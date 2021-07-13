package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// gin的helloWorld

func main() {
	// 1.创建路由
	r := gin.Default()
	// 2.绑定路由规则
	// gin.Context，封装了request和response
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello World!")
	})
	r.POST()
	// 3.监听端口，默认再8080
	r.Run(":8000")
}
