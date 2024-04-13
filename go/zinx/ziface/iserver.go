package ziface

// IServer 定义服务器接口
type IServer interface {
	// Start 启动服务器
	Start()
	// Stop 停止服务器
	Stop()
	// Serve 开启业务服务
	Serve()
	// AddRouter 添加路由
	AddRouter(msgId uint32, router IRouter)

	// GetConnMgr 获取连接管理器
	GetConnMgr() IConnManager

	SetOnConnStart(func(IConnection))
	SetOnConnStop(func(IConnection))

	CallOnConnStart(conn IConnection)
	CallOnConnStop(conn IConnection)
}
