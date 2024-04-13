package config

import (
	"github.com/unknwon/goconfig"
	"os"
)

const configFile = "/conf/conf.ini"

var File *goconfig.ConfigFile

func init() {
	currentDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	configPath := currentDir + configFile
	if len(os.Args) > 1 {
		dir := os.Args[1]
		configPath = dir + configFile
	}
	if fileExist(configPath) != nil {
		panic("配置文件不存在")
	}
	File, err = goconfig.LoadConfigFile(configPath)
	if err != nil {
		panic(err)
	}
}
func fileExist(configPath string) error {
	_, err := os.Stat(configPath)
	return err
}
