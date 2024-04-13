package main

import (
	"crawler/collect"
	"crawler/log"
	"crawler/proxy"
	"fmt"
	"go.uber.org/zap/zapcore"
	"time"
)

func main() {
	plugin, c := log.NewFilePlugin("./log.txt", zapcore.InfoLevel)
	defer c.Close()
	logger := log.NewLogger(plugin)
	logger.Info("greeter world")

	url := "https://google.com"
	// export http_proxy=http://127.0.0.1:15732 https_proxy=http://127.0.0.1:15732
	proxyUrls := []string{"http://127.0.0.1:15732"}
	p, err := proxy.RoundRobinProxySwitcher(proxyUrls...)
	f := collect.BrowserFetch{
		Timeout: 3000 * time.Millisecond,
		Proxy:   p,
	}
	body, err := f.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf(string(body))

}
