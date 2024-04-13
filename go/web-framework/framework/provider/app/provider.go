package app

import (
	"web-framework/framework"
	"web-framework/framework/contract"
)

type WebAppProvider struct {
	BaseFolder string
}

func (w *WebAppProvider) Register(container framework.Container) framework.NewInstance {
	return NewWebApp
}

func (w *WebAppProvider) Boot(container framework.Container) error {
	return nil
}

func (w *WebAppProvider) Params(container framework.Container) []interface{} {
	return []interface{}{container, w.BaseFolder}
}

func (w *WebAppProvider) Name() string {
	return contract.AppKey
}

func (w *WebAppProvider) IsDefer() bool {
	return false
}
