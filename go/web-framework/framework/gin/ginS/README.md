# Gin Default Server

This is API experiment for Gin.

```go
package main

import (
	"web-framework/framework/gin"
	"web-framework/framework/gin/ginS"
)

func main() {
	ginS.GET("/", func(c *gin.Context) { c.String(200, "Hello World") })
	ginS.Run()
}
```
