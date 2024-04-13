package main

import (
	"fmt"
	"three-kingdoms/config"
)

func main() {
	_ = config.File.MustValue("login_server", "host", "127.0.0.1")
	fmt.Println("start success")
}
