package framework

type NewInstance func(...interface{}) (interface{}, error)
type ServiceProvider interface {
	Register(Container) NewInstance
	Boot(Container) error
	Params(Container) []interface{}
	Name() string
	IsDefer() bool
}
