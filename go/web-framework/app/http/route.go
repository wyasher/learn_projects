package http

import (
	"web-framework/app/http/module/demo"
	"web-framework/framework/gin"
)

func Route(r *gin.Engine) {
	r.Static("/dist/", "./dist/")
	demo.Register(r)
}
