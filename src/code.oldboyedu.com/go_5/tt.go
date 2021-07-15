package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

// 定义中间件
//func MiddleWare() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		t := time.Now()
//		fmt.Println("中间件开始执行")
//		// 设置变量到context的key中
//		c.Set("request", "中间件")
//		// 执行中间件
//		c.Next()
//		status := c.Writer.Status()
//		fmt.Println("中间件执行完成", status)
//		t2 := time.Since(t)
//		fmt.Println("time", t2)
//	}
//}

func myTime(c *gin.Context) {
	start := time.Now()
	c.Next()
	since := time.Since(start)
	fmt.Println("程序用时:", since)
}

func main() {
	// 1.创建路由
	r := gin.Default()
	// 注册中间件
	//r.Use(myTime)
	shoppingGroup := r.Group("/shopping")
	shoppingGroup.Use(myTime)
	{
		shoppingGroup.GET("/index", shopIndexHandler)
		shoppingGroup.GET("/home", shopHomeHandler)
	}

	r.Run(":8000")
}

func shopIndexHandler(c *gin.Context) {
	time.Sleep(5 * time.Second)
}

func shopHomeHandler(c *gin.Context) {
	time.Sleep(3 * time.Second)
}
