package main

import (
	"web-framework/app/console"
	"web-framework/app/http"
	"web-framework/framework"
	_ "web-framework/framework/middleware"
	"web-framework/framework/provider/app"
	"web-framework/framework/provider/kernel"
)

func main() {
	container := framework.NewWebContainer()
	container.Bind(&app.WebAppProvider{})
	if engine, err := http.NewHttpEngine(); err == nil {
		container.Bind(&kernel.WebKernelProvider{
			HttpEngine: engine,
		})
	}
	console.RunCommand(container)
}
