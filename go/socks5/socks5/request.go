package socks5

import (
	"context"
	"fmt"
	"io"
	"net"
	"strings"
)

// 请求命令
const (
	ConnectCommand      uint8 = 0x01 // CONNECT请求
	BindCommand         uint8 = 2    //表示BIND请求
	UDPAssociateCommand uint8 = 3    //UDP传输请求
)

// 回复命令
const (
	succeeded               uint8 = iota // 成功
	serverFailure                        // 服务器失败
	rulesetNotAll                        //  规则集不允许连接
	networkUnreachable                   // 网络不可达
	hostUnreachable                      // 主机不可达
	connectionRefused                    // 连接拒绝
	ttlExpired                           // ttl过期
	commandNotSupported                  // 命令不支持
	addressTypeNotSupported              // 地址类型不支持
)

// 保留字
const rsvValue uint8 = 0x00

// Request 客户端请求信息
//
//	+----+-----+-------+------+----------+----------+
//	|VER | CMD |  RSV  | ATYP | DST.ADDR | DST.PORT |
//	+----+-----+-------+------+----------+----------+
//	| 1  |  1  | X'00' |  1   | Variable |    2     |
//	+----+-----+-------+------+----------+----------+
//
// VER 版本号，socks5的值为0x05
// CMD 命令：0x01表示CONNECT请求 0x02表示BIND请求 0x03表示UDP转发
// RSV 保留字段，值为0x00
// ATYP 目标地址类型，DST.ADDR的数据对应这个字段的类型。
//   - 0x01表示IPv4地址，DST.ADDR为4个字节
//   - 0x03表示域名，DST.ADDR是一个可变长度的域名
//   - 表示IPv6地址，DST.ADDR为16个字节长度
//
// DST.ADDR 一个可变长度的值
// DST.PORT 目标端口，固定2个字节
type Request struct {
	// socks 版本
	Version uint8
	// 请求命令
	Command uint8
	// 认证信息
	AuthContext *AuthContext
	// 请求客户地址信息
	RemoteAddr AddressSpec
	// 目的地地址信息
	DestAddr AddressSpec
	// 实际目标（可能受重写影响）
	realDestAddr AddressSpec
	bufConn      io.Reader
}

// NewRequest 通过连接创建一个请求
func NewRequest(bufConn io.Reader) (*Request, error) {
	// 读取 version command rsv  rsv忽略不做使用
	header := []byte{0, 0, 0}
	if _, err := io.ReadAtLeast(bufConn, header, len(header)); err != nil {
		return nil, fmt.Errorf("failed to get request header info: %v", err)
	}
	// 确保时socks5
	if header[0] != SocksVersion5 {
		return nil, fmt.Errorf("unsupported command version: %d", header[0])
	}
	// 读取目标地址
	dest, err := NewAddressSpec(bufConn)
	if err != nil {
		return nil, err
	}

	request := &Request{
		Version:  SocksVersion5,
		Command:  header[1],
		DestAddr: dest,
		bufConn:  bufConn,
	}

	return request, err
}

// conn 隐藏一些net.Conn的其他实现
type conn interface {
	Write([]byte) (int, error)
	RemoteAddr() net.Addr
}

// handleBind 用于处理bind command
func (s *Server) handleBind(ctx context.Context, conn conn, req *Request) error {
	// TODO 暂时不支持bind command
	if err := sendReply(conn, commandNotSupported, nil); err != nil {
		return err
	}
	return nil
}

// handleAssociate 用于处理 udp associate command
func (s *Server) handleUDPAssociate(ctx context.Context, conn conn, req *Request) error {
	// TODO 暂时不支持udp associate command
	if err := sendReply(conn, commandNotSupported, nil); err != nil {
		return err
	}
	return nil
}

// handleConnect 处理 connect command
func (s *Server) handleConnect(ctx context.Context, conn conn, req *Request) error {
	// 尝试连接
	dial := s.config.Dial
	if dial == nil {
		dial = func(ctx context.Context, network, addr string) (net.Conn, error) {
			return net.Dial(network, addr)
		}
	}

	target, err := dial(ctx, "tcp", req.realDestAddr.Address())
	if err != nil {
		msg := err.Error()
		resp := hostUnreachable
		if strings.Contains(msg, "refused") {
			resp = connectionRefused
		} else if strings.Contains(msg, "network is unreachable") {
			resp = networkUnreachable
		}
		if err := sendReply(conn, resp, nil); err != nil {
			return err
		}
		return fmt.Errorf("connect to %v failed: %v", req.DestAddr, err)
	}
	defer func(target net.Conn) {
		_ = target.Close()
	}(target)
	// 回复连接成功
	local := target.LocalAddr().(*net.TCPAddr)
	bind := ConvertTcpAddr(local)
	if err := sendReply(conn, succeeded, bind); err != nil {
		return err
	}
	// 开启代理
	errCh := make(chan error, 2)
	go proxy(target, req.bufConn, errCh)
	go proxy(conn, target, errCh)
	for i := 0; i < 2; i++ {
		e := <-errCh
		if e != nil {
			// return from this function closes target (and conn).
			return e
		}
	}
	return nil
}

// handleRequest 处理客户端请求
func (s *Server) handleRequest(req *Request, conn conn) error {
	ctx := context.Background()

	dest := req.DestAddr
	fqdnAddr, ok := dest.(*FQDNAddrSpec)
	// 如果是域名 解析解析域名
	if ok {
		ctx_, addr, err := s.config.Resolver.Resolve(ctx, fqdnAddr.FQDN)
		if err != nil {
			if err := sendReply(conn, hostUnreachable, nil); err != nil {
				return err
			}
			return fmt.Errorf("failed to resolve destination '%s': %v", fqdnAddr.FQDN, err)
		}
		ctx = ctx_
		fqdnAddr.IP = addr
	}
	req.realDestAddr = req.DestAddr
	// 地址重写
	if s.config.Rewriter != nil {
		ctx, req.realDestAddr = s.config.Rewriter.Rewrite(ctx, req)
	}
	// 检查是否连接是否被允许
	if ctx_, ok := s.config.Rules.Allow(ctx, req); !ok {
		if err := sendReply(conn, rulesetNotAll, nil); err != nil {
			return err
		}
		return fmt.Errorf("bind to %v blocked by rules", req.DestAddr)
	} else {
		ctx = ctx_
	}
	switch req.Command {
	case ConnectCommand:
		return s.handleConnect(ctx, conn, req)
	case BindCommand:
		return s.handleBind(ctx, conn, req)
	case UDPAssociateCommand:
		return s.handleUDPAssociate(ctx, conn, req)
	default:
		if err := sendReply(conn, commandNotSupported, nil); err != nil {
			return err
		}
		return fmt.Errorf("unsupported command: %d", req.Command)
	}
}

type closeWriter interface {
	CloseWrite() error
}

// proxy 将src数据发送到dst
func proxy(dst io.Writer, src io.Reader, errCh chan error) {
	_, err := io.Copy(dst, src)
	if tcpConn, ok := dst.(closeWriter); ok {
		_ = tcpConn.CloseWrite()
	}
	errCh <- err
}
