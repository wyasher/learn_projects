package kernel

import (
	"net/http"
	"web-framework/framework/gin"
)

type WebKernelService struct {
	engine *gin.Engine
}

func NewWebKernelService(params ...interface{}) (interface{}, error) {
	return &WebKernelService{
		engine: params[0].(*gin.Engine),
	}, nil
}
func (s *WebKernelService) HttpEngine() http.Handler {
	return s.engine
}
