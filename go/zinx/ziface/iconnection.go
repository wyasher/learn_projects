package ziface

import "net"

type IConnection interface {
	// Start 启动连接，让当前连接开始工作
	Start()
	// Stop 停止连接，结束当前连接状态M
	Stop()
	// GetTCPConnection 获取当前连接绑定的socket conn
	GetTCPConnection() *net.TCPConn
	// GetConnID 获取当前连接模块的连接ID
	GetConnID() uint32
	// RemoteAddr 获取远程客户端地址信息
	RemoteAddr() net.Addr
	// SendMsg 发送数据，将数据发送给远程的客户端
	SendMsg(msgId uint32, data []byte) error

	SendBuffMsg(msgId uint32, data []byte) error

	SetProperty(key string, value any)

	GetProperty(key string) (any, error)

	RemoveProperty(key string)
}

// HandFunc 定义一个处理连接业务的方法
type HandFunc func(*net.TCPConn, []byte, int) error
