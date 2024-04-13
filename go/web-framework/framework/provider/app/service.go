package app

import (
	"errors"
	"flag"
	"path/filepath"
	"web-framework/framework"
	"web-framework/framework/util"
)

type WebApp struct {
	container  framework.Container
	baseFolder string
}

func (w WebApp) Version() string {
	return "0.0.3"
}

func (w WebApp) BaseFolder() string {
	if w.baseFolder != "" {
		return w.baseFolder
	}
	var baseFolder string
	flag.StringVar(&baseFolder, "base_folder", "", "base_folder参数，默认当前路径")
	flag.Parsed()
	if baseFolder != "" {
		return baseFolder
	}
	return util.GetExecDirectory()
}

func (w WebApp) ConfigFolder() string {
	return filepath.Join(w.baseFolder, "config")
}
func (w WebApp) StorageFolder() string {
	return filepath.Join(w.baseFolder, "storage")
}

func (w WebApp) LogFolder() string {
	return filepath.Join(w.StorageFolder(), "log")
}

func (w WebApp) ProviderFolder() string {
	return filepath.Join(w.baseFolder, "provider")
}

func (w WebApp) MiddlewareFolder() string {
	return filepath.Join(w.baseFolder, "middleware")

}

func (w WebApp) CommandFolder() string {
	return filepath.Join(w.baseFolder, "command")
}

func (w WebApp) RuntimeFolder() string {
	return filepath.Join(w.baseFolder, "runtime")
}

func (w WebApp) TestFolder() string {
	return filepath.Join(w.baseFolder, "test")
}

func NewWebApp(params ...interface{}) (interface{}, error) {
	if len(params) != 2 {
		return nil, errors.New("params error")
	}
	container := params[0].(framework.Container)
	baseFolder := params[1].(string)
	return &WebApp{
		baseFolder: baseFolder,
		container:  container,
	}, nil
}
