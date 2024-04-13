package ziface

type IRouter interface {
	PreHandle(request IRequest)  // 处理业务之前的方法
	Handle(request IRequest)     // 处理业务的方法
	PostHandle(request IRequest) // 处理业务之后的方法
}
