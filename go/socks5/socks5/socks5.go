package socks5

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

const (
	SocksVersion5 = uint8(5) // socks5版本号
)

// Config SOCKS5 服务配置
type Config struct {
	// AuthMethods 提供认证方法实现 默认不需要认证
	AuthMethods []Authenticator
	// 当使用UserPassAuthenticator需要提供
	Credentials CredentialStore
	// 日志
	Logger *log.Logger
	// 提供自定义域名解析 默认使用DNSResolver
	Resolver NameResolver
	// 提供自定义规则 默认允许所有 PermitAll
	Rules RuleSet

	Rewriter AddressRewriter

	BindIP net.IP

	Dial func(ctx context.Context, network, addr string) (net.Conn, error)
}

// Server 负责接受连接和处理SOCKS5协议
type Server struct {
	// 服务配置
	config *Config
	// 验证器映射
	authenticators map[uint8]Authenticator
}

// New 创建一个SOCKS5服务
func New(conf *Config) (*Server, error) {
	// 设置默认认证器
	if len(conf.AuthMethods) == 0 {
		if conf.Credentials != nil {
			// Credentials 不为nil时需要使用UserPassAuthenticator验证器
			conf.AuthMethods = []Authenticator{&UserPassAuthenticator{conf.Credentials}}
		} else {
			conf.AuthMethods = []Authenticator{&NoAuthAuthenticator{}}
		}
	}
	// 默认时用DNSResolver
	if conf.Resolver == nil {
		conf.Resolver = DNSResolver{}
	}
	// 设置默认规则
	if conf.Rules == nil {
		conf.Rules = DefaultRuleSet()
	}
	// 设置默认logger
	if conf.Logger == nil {
		conf.Logger = log.New(os.Stdout, "", log.LstdFlags)
	}

	server := &Server{
		config: conf,
	}
	// 设置支持认证 map
	server.authenticators = make(map[uint8]Authenticator)
	for _, a := range conf.AuthMethods {
		server.authenticators[a.GetAuthMethod()] = a
	}
	return server, nil

}

// ListenAndServe 开启服务
func (s *Server) ListenAndServe(network, addr string) error {
	l, err := net.Listen(network, addr)
	if err != nil {
		return err
	}
	return s.Serve(l)

}

// Serve 监听连接
func (s *Server) Serve(l net.Listener) error {
	for {
		conn, err := l.Accept()
		if err != nil {
			return err
		}
		go func() {
			_ = s.ServeConn(conn)
		}()
	}
}

// ServeConn 处理连接
func (s *Server) ServeConn(conn net.Conn) error {
	defer func(conn net.Conn) {
		_ = conn.Close()
	}(conn)

	bufConn := bufio.NewReader(conn)

	// 读取socks version 并验证是否支持
	socksVersion := []byte{0}
	if _, err := bufConn.Read(socksVersion); err != nil {
		s.config.Logger.Printf("[ERR] socks: Failed to get version byte: %v", err)
		return err
	}
	if socksVersion[0] != SocksVersion5 {
		err := fmt.Errorf("unsupported SOCKS version: %d", socksVersion[0])
		s.config.Logger.Printf("[ERR] socks: %v", err)
		return err
	}
	// 认证连接
	authContext, err := s.authenticate(conn, bufConn)
	if err != nil {
		err = fmt.Errorf("failed to authenticate: %v", err)
		s.config.Logger.Printf("[ERR] socks: %v", err)
		return err
	}
	// 创建请求
	request, err := NewRequest(bufConn)
	if err != nil {
		if errors.Is(err, unrecognizedAddrType) {
			// 发送不支持的回复
			if err := sendReply(conn, addressTypeNotSupported, nil); err != nil {
				return err
			}
		}
		return fmt.Errorf("failed to create request: %v", err)
	}
	request.AuthContext = authContext
	if client, ok := conn.RemoteAddr().(*net.TCPAddr); ok {
		request.RemoteAddr = ConvertTcpAddr(client)
	}

	// 处理请求
	if err := s.handleRequest(request, conn); err != nil {
		err = fmt.Errorf("failed to handle request: %v", err)
		s.config.Logger.Printf("[ERR] socks: %v", err)
		return err
	}

	return nil

}

// sendReply 发送回复
//
//	+----+-----+-------+------+----------+----------+
//	|VER | REP |  RSV  | ATYP | BND.ADDR | BND.PORT |
//	+----+-----+-------+------+----------+----------+
//	| 1  |  1  | X'00' |  1   | Variable |    2     |
//	+----+-----+-------+------+----------+----------+
//
// VER socks版本，这里为0x05
// REP Relay field,内容取值如下
//   - X’00’ succeeded
//   - X’01’ general SOCKS server failure
//   - X’02’ connection not allowed by ruleset
//   - X’03’ Network unreachable
//   - X’04’ Host unreachable
//   - X’05’ Connection refused
//   - X’06’ TTL expired
//   - X’07’ Command not supported
//   - X’08’ Address type not supported
//   - X’09’ to X’FF’ unassigned
//
// RSV 保留字段
// ATYPE 同请求的ATYPE
// BND.ADDR 服务绑定的地址
// BND.PORT 服务绑定的端口DST.PORT
func sendReply(w io.Writer, respCommand uint8, addr AddressSpec) error {
	// 3 为 VER + REP + RSV 长度
	reply := make([]byte, 3+addr.Len())
	// VER
	reply[0] = SocksVersion5
	// REP
	reply[1] = respCommand
	// RSV
	reply[2] = rsvValue
	// encode 地址部分数据
	addrBytes := addr.Encode()
	copy(reply[3:], addrBytes)
	// 发送数据
	_, err := w.Write(reply)
	if err != nil {
		return fmt.Errorf("failed to send reply: %v", err)
	}
	return err
}
