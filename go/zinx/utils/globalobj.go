package utils

import (
	"encoding/json"
	"os"
	"zinx/ziface"
)

type GlobalObj struct {
	TcpServer        ziface.IServer // 当前全局的server对象
	Host             string         // 当前服务器主机监听的IP
	TcpPort          int            // 当前服务器主机监听的端口号
	Name             string         // 当前服务器的名称
	Version          string         // 当前Zinx版本号
	MaxConn          int            // 当前服务器主机允许的最大连接数
	MaxPackageSize   uint32         // 当前Zinx框架数据包的最大值
	WorkerPoolSize   uint32         // 当前业务工作Worker池的Goroutine数量
	MaxWorkerTaskLen uint32         // woker 对应负责的任务队列的最大任务存储数量
	ConfFilePath     string         // 配置文件路径
	MaxMsgChanLen    uint32
}

// GlobalObject 全局的对外GlobalObj
var GlobalObject *GlobalObj

// Reload 从zinx.json去加载自定义的参数
func (g *GlobalObj) Reload() {
	data, err := os.ReadFile("conf/zinx.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, &GlobalObject)
	if err != nil {
		panic(err)
	}
}
func init() {
	GlobalObject = &GlobalObj{
		Name:             "ZinxServerApp",
		Version:          "V0.10",
		TcpPort:          7777,
		Host:             "0.0.0.0",
		MaxConn:          12000,
		MaxPackageSize:   4096,
		WorkerPoolSize:   10,
		ConfFilePath:     "conf/zinx.json",
		MaxWorkerTaskLen: 1024,
		MaxMsgChanLen:    1024,
	}
	GlobalObject.Reload()
}
