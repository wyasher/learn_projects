package framework

import (
	"errors"
	"fmt"
	"sync"
)

type Container interface {
	Bind(provider ServiceProvider) error
	IsBind(key string) bool
	MustMake(key string) interface{}
	Make(key string) (interface{}, error)
	MakeNew(key string, params ...interface{}) (interface{}, error)
}
type WebContainer struct {
	Container
	providers map[string]ServiceProvider
	instances map[string]interface{}
	lock      sync.RWMutex
}

func (web *WebContainer) Bind(provider ServiceProvider) error {
	web.lock.Lock()
	defer web.lock.Unlock()
	key := provider.Name()
	web.providers[key] = provider
	// 如果不是延迟初始化 就立即初始化
	if !provider.IsDefer() {
		if err := provider.Boot(web); err != nil {
			return err
		}
		// 初始化
		params := provider.Params(web)
		method := provider.Register(web)
		instance, err := method(params...)
		if err != nil {
			return errors.New(err.Error())
		}
		web.instances[key] = instance
	}
	return nil
}

func (web *WebContainer) IsBind(key string) bool {
	web.lock.RLock()
	defer web.lock.RUnlock()
	_, ok := web.providers[key]
	return ok
}

func (web *WebContainer) MustMake(key string) interface{} {
	instance, err := web.make(key, nil, true)
	if err != nil {
		panic(err)
	}
	return instance
}
func (web *WebContainer) MakeNew(key string, params ...interface{}) (interface{}, error) {
	return web.make(key, params, true)
}
func (web *WebContainer) Make(key string) (interface{}, error) {
	return web.make(key, nil, false)
}

func (web *WebContainer) make(key string, params []interface{}, forceNew bool) (interface{}, error) {
	web.lock.RLock()
	defer web.lock.RUnlock()
	sp := web.findServiceProvider(key)
	if sp == nil {
		return nil, errors.New("contract " + key + " have not register")
	}
	if forceNew {
		return web.newInstance(sp, params)
	}
	if instance, ok := web.instances[key]; ok {
		return instance, nil
	}
	// 实例化
	instance, err := web.newInstance(sp, params)
	if err != nil {
		return nil, err
	}
	web.instances[key] = instance
	return instance, nil
}

func NewWebContainer() *WebContainer {
	return &WebContainer{
		providers: map[string]ServiceProvider{},
		instances: map[string]interface{}{},
		lock:      sync.RWMutex{},
	}
}

func (web *WebContainer) PrintProviders() []string {
	var ret []string
	for _, provider := range web.providers {
		name := provider.Name()
		line := fmt.Sprint(name)
		ret = append(ret, line)
	}
	return ret
}
func (web *WebContainer) findServiceProvider(key string) ServiceProvider {
	web.lock.RLock()
	defer web.lock.RUnlock()
	if sp, ok := web.providers[key]; ok {
		return sp
	}
	return nil
}

func (web *WebContainer) newInstance(sp ServiceProvider, params []interface{}) (interface{}, error) {
	if err := sp.Boot(web); err != nil {
		return nil, err
	}
	if params == nil {
		params = sp.Params(web)
	}
	method := sp.Register(web)
	instance, err := method(params...)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return instance, nil
}
