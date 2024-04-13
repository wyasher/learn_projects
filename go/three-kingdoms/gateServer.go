package main

import (
	"three-kingdoms/config"
	"three-kingdoms/net"
	"three-kingdoms/server/gate"
)

func main() {
	host := config.File.MustValue("gate_server", "host", "127.0.0.1")
	port := config.File.MustValue("gate_server", "port", "8004")
	s := net.NewServer(host + ":" + port)
	s.NeedSecret(true)
	gate.Init()
	s.Router(gate.Router)
	s.Start()
}
