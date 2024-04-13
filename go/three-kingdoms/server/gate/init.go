package gate

import (
	"three-kingdoms/net"
	"three-kingdoms/server/gate/controller"
)

var Router = &net.Router{}

func Init() {
	initRouter()
}

func initRouter() {
	controller.GateHandler.Router(Router)
}
