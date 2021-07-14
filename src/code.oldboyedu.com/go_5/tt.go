package main

import (
<<<<<<< HEAD
	"fmt"
	"github.com/gin-gonic/gin"
=======
	"github.com/gin-gonic/gin"
	"net/http"
>>>>>>> origin/master
)

// gin的helloWorld

func main() {
	// 1.创建路由
<<<<<<< HEAD
	// 默认使用了2个中间件logger(), Recovery()
	r := gin.Default()
	// 路由组1,处理GET请求
	v1 := r.Group("/v1")
	{
		v1.GET("/login", login)
		v1.GET("/submit", submit)
	}
	v2 := r.Group("/v2")
	{
		v2.POST("/login", login)
		v2.POST("/submit", submit)
	}
	r.Run(":8000")
}

func login(c *gin.Context) {
	name := c.DefaultQuery("name", "jack")
	c.String(200, fmt.Sprintf("hello %s\n", name))
}

func submit(c *gin.Context) {
	name := c.DefaultQuery("name", "lily")
	c.String(200, fmt.Sprintf("hello %s\n", name))
}
=======
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
>>>>>>> origin/master
