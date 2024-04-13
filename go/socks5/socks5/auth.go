package socks5

import (
	"errors"
	"fmt"
	"io"
)

// 认证协议
// +----+----------+----------+
// |VER | NMETHODS | METHODS  |
// +----+----------+----------+
// | 1  |    1     | 1 to 255 |
// +----+----------+----------+
// VER: 协议版本，socks5为0x05
// NMETHODS: 客户端支持认证的方法数量
// METHODS: 可用的认证方式列表
//	X’00’ NO AUTHENTICATION REQUIRED
//	X’01’ GSSAPI
//	X’02’ USERNAME/PASSWORD
//	X’03’ to X’7F’ IANA ASSIGNED
//	X’80’ to X’FE’ RESERVED FOR PRIVATE METHODS
//	X’FF’ NO ACCEPTABLE METHODS
// 认证结果协议
//	+----+--------+
//	|VER | STATUS |
//	+----+--------+
//	| 1  |   1    |
//	+----+--------+
// VER:认证version
// STATUS:认证状态 0x00 success 其他 失败

const (
	NoAuth             uint8 = 0x00 // 不需要认证
	GSSAPI                   = 0x01 // GSSAPI 未实现
	UsernamePassword         = 0x02 // 账号密码方式
	IanaAssigned             = 0x03
	ReservedForPrivate       = 0x80 // 私有
	NoAcceptable             = 0xFF // 不支持方式 客户端需要断开链接
)

const (
	usernamePasswordVersion1 uint8 = 0x01 // 用户密码认证支持的版本
	authSuccess                    = 0x00 // 认证成功
	authFail                       = 0x01 // 认证失败
)

var (
	UserAuthFailed  = errors.New("user authentication failed")
	NoSupportedAuth = errors.New("no supported authentication mechanism")
)

// AuthContext 认证上下文
type AuthContext struct {
	// 认证方式
	Method uint8
	// 认证上下文数据 比如账号信息之类
	Payload map[string]string
}

// Authenticator 身份验证器
type Authenticator interface {
	// Authenticate 认证
	// 认证需要返回协议格式
	//	+----+--------+
	//	|VER | METHOD |
	//	+----+--------+
	//	| 1  |   1    |
	//	+----+--------+
	Authenticate(reader io.Reader, writer io.Writer) (*AuthContext, error)
	// GetAuthMethod 获取认证方式
	GetAuthMethod() uint8
}

// NoAuthAuthenticator 不需要认证的认证器
type NoAuthAuthenticator struct{}

func (n *NoAuthAuthenticator) Authenticate(_ io.Reader, writer io.Writer) (*AuthContext, error) {
	// 告诉客户端不需要认证
	_, err := writer.Write([]byte{SocksVersion5, byte(NoAuth)})
	return &AuthContext{
		Method:  NoAuth,
		Payload: nil,
	}, err
}

func (n *NoAuthAuthenticator) GetAuthMethod() uint8 {
	return NoAuth
}

// UserPassAuthenticator 用户账号密码的认证器
type UserPassAuthenticator struct {
	// 账号密码校验
	Credentials CredentialStore
}

// Authenticate 账号密码认证
// 认证协议
//
//	+----+------+----------+------+----------+
//	|VER | ULEN |  UNAME   | PLEN |  PASSWD  |
//	+----+------+----------+------+----------+
//	| 1  |  1   | 1 to 255 |  1   | 1 to 255 |
//	+----+------+----------+------+----------+
//	VER: 版本号 目前支持 0x01
//	ULEN: 用户名长度
//	UNAME: 对应用户名的字节数据
//	PLEN: 密码长度
//	PASSWD: 密码对应的数据
func (u *UserPassAuthenticator) Authenticate(reader io.Reader, writer io.Writer) (*AuthContext, error) {
	// 告诉客户端使用账号密码认证方式
	if _, err := writer.Write([]byte{SocksVersion5, byte(UsernamePassword)}); err != nil {
		return nil, err
	}

	// 获取socks版本和用户名长度
	header := []byte{0, 0}
	if _, err := io.ReadAtLeast(reader, header, len(header)); err != nil {
		return nil, err
	}
	// 验证是否是支持的认证版本号
	if header[0] != usernamePasswordVersion1 {
		return nil, fmt.Errorf("auth version not support:%v\n", header[0])
	}
	// 获取用户名称
	usernameLen := header[1]
	username := make([]byte, usernameLen)
	if _, err := io.ReadAtLeast(reader, username, int(usernameLen)); err != nil {
		return nil, err
	}
	// 获取用户密码
	if _, err := reader.Read(header[:1]); err != nil {
		return nil, err
	}
	passwordLen := header[0]
	password := make([]byte, passwordLen)
	if _, err := io.ReadAtLeast(reader, password, int(passwordLen)); err != nil {
		return nil, err
	}
	// 验证账号密码是否正确
	if u.Credentials.Valid(string(username), string(password)) {
		if _, err := writer.Write([]byte{usernamePasswordVersion1, authSuccess}); err != nil {
			return nil, err
		}
	} else {
		// 认证失败
		if _, err := writer.Write([]byte{usernamePasswordVersion1, authFail}); err != nil {
			return nil, err
		}
		return nil, UserAuthFailed
	}
	// 认证成功保存
	return &AuthContext{
		UsernamePassword,
		map[string]string{
			"Username": string(username),
		},
	}, nil

}

func (u *UserPassAuthenticator) GetAuthMethod() uint8 {
	return UsernamePassword
}

// authenticate 验证连接
func (s *Server) authenticate(conn io.Writer, bufConn io.Reader) (*AuthContext, error) {
	// 获取methods
	methods, err := readMethods(bufConn)
	if err != nil {
		return nil, fmt.Errorf("failed to get auth methods: %v", err)
	}
	// 选择一个验证器 执行验证
	// 使用最先匹配上的
	for _, method := range methods {
		authenticator, ok := s.authenticators[method]
		if ok {
			return authenticator.Authenticate(bufConn, conn)
		}
	}
	return nil, noAcceptableAuth(conn)
}

// noAcceptableAuth 没有适配的认证方法
func noAcceptableAuth(conn io.Writer) error {
	_, _ = conn.Write([]byte{SocksVersion5, NoAcceptable})
	return NoSupportedAuth
}

// readMethods 获取认证方法（METHODS）
// 认证协议
// +----+----------+----------+
// |VER | NMETHODS | METHODS  |
// +----+----------+----------+
// | 1  |    1     | 1 to 255 |
// +----+----------+----------+
func readMethods(r io.Reader) ([]byte, error) {
	// 读取支持的METHODS数量 NMETHODS
	header := []byte{0}
	if _, err := r.Read(header); err != nil {
		return nil, err
	}
	numMethods := int(header[0])
	// METHODS
	methods := make([]byte, numMethods)
	_, err := io.ReadAtLeast(r, methods, numMethods)
	return methods, err
}
