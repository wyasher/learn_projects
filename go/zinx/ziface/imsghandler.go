package ziface

type IMsgHandle interface {
	DoMsgHandler(request IRequest)
	AddRouter(msgId uint32, router IRouter)
	StartWorkerPool()                    // 启动worker工作池
	SendMsgToTaskQueue(request IRequest) // 将消息交给TaskQueue，由worker进行处理
}
