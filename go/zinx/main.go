package main

import (
	"fmt"
	"zinx/ziface"
	"zinx/znet"
)

type PingRouter struct {
	znet.BaseRouter
}

func (p *PingRouter) Handle(request ziface.IRequest) {
	fmt.Println("Call Router Handle...")
	fmt.Println("recv from client: msgID=", request.GetMsgID(), ", data=", string(request.GetData()))
	err := request.GetConnection().SendMsg(1, []byte("ping...ping...ping"))
	if err != nil {
		fmt.Println(err)
	}
}

type HelloZinxRouter struct {
	znet.BaseRouter
}

func (h *HelloZinxRouter) Handle(request ziface.IRequest) {
	fmt.Println("Call HelloZinxRouter Handle...")
	fmt.Println("recv from client: msgID=", request.GetMsgID(), ", data=", string(request.GetData()))
	err := request.GetConnection().SendMsg(201, []byte("Hello Zinx Router V0.6"))
	if err != nil {
		fmt.Println(err)
	}
}

func DoConnectingBegin(conn ziface.IConnection) {
	fmt.Println("DoConnectingBegin is Called...")
	conn.SetProperty("Name", "Aceld")
	conn.SetProperty("Home", "https://github.com/aceld/zinx")
	err := conn.SendMsg(2, []byte("DoConnectingBegin is Called..."))
	if err != nil {
		fmt.Println(err)
	}
}
func DoConnectLost(conn ziface.IConnection) {
	fmt.Println("DoConnectLost is Called...")
	if name, err := conn.GetProperty("Name"); err == nil {
		fmt.Println("Name = ", name)
	}
	if home, err := conn.GetProperty("Home"); err == nil {
		fmt.Println("Home = ", home)
	}
}

func main() {
	s := znet.NewServer()
	s.SetOnConnStop(DoConnectLost)
	s.SetOnConnStart(DoConnectingBegin)
	s.AddRouter(0, &PingRouter{})
	s.AddRouter(1, &HelloZinxRouter{})
	s.Serve()
}
