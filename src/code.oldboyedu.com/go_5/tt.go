package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
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

func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取客户端cookie并校验
		if cookie, err := c.Cookie("abc"); err == nil {
			if cookie == "123" {
				c.Next()
				return
			}
		}
		//  返回错误
		c.JSON(http.StatusUnauthorized, gin.H{"error": "err"})
		// 若验证不通过，不再调用后续的函数处理
		c.Abort()
		return
	}
}

func main() {
	// 1.创建路由
	r := gin.Default()
	r.GET("/login", func(c *gin.Context) {
		// 设置cookie
		c.SetCookie("abc", "123", 60, "/",
			"localhost", false, true)
		// 返回信息
		c.String(200, "Login success!")
	})
	r.GET("/home", AuthMiddleWare(), func(c *gin.Context) {
		c.JSON(200, gin.H{"data": "home"})
	})
	r.Run(":8000")
}
