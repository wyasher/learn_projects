package ziface

type IRequest interface {
	GetConnection() IConnection // 获取请求连接信息
	GetData() []byte
	GetMsgID() uint32 // 获取请求消息的数据
}
