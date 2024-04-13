package login

import (
	"three-kingdoms/db"
	"three-kingdoms/net"
	"three-kingdoms/server/login/controller"
)

var Router = net.NewRouter()

func Init() {
	db.TestDB()
	initRouter()
}

func initRouter() {
	controller.DefaultAccount.Router(Router)
}
