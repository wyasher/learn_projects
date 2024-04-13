package kernel

import (
	"web-framework/framework"
	"web-framework/framework/contract"
	"web-framework/framework/gin"
)

type WebKernelProvider struct {
	HttpEngine *gin.Engine
}

func (provider *WebKernelProvider) Register(c framework.Container) framework.NewInstance {
	return NewWebKernelService
}

func (provider *WebKernelProvider) Boot(c framework.Container) error {
	if provider.HttpEngine == nil {
		provider.HttpEngine = gin.Default()
	}
	provider.HttpEngine.SetContainer(c)
	return nil
}

func (provider *WebKernelProvider) Params(c framework.Container) []interface{} {
	return []interface{}{provider.HttpEngine}
}
func (provider *WebKernelProvider) Name() string {
	return contract.KernelKey
}
func (provider *WebKernelProvider) IsDefer() bool {
	return false
}
