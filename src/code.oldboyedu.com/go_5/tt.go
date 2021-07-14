package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 定义接受数据的结构体
type Login struct {
	// binding:"required" 修饰的字段，若接受为空值，则报错
	User     string `form:"username" json:"user" uri:"user" xml:"user" binding:"required"`
	Password string `form:"password" json:"password" uri:"password" xml:"password" binding:"required"`
}

func main() {
	// 1.创建路由
	r := gin.Default()
	// JSON绑定
	r.GET("/:user/:password", func(c *gin.Context) {
		// 声明接收的变量
		var login Login
		// Bind()默认解析并绑定form格式
		// 根据请求头中content-type自动推断
		if err := c.ShouldBindUri(&login); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// JSON绑定
		// 将request的body种的数据，自动按照json格式解析到结构体
		//if err := c.ShouldBindJSON(&json); err != nil {
		//	c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		//	return
		//}
		// 判断用户名密码是否正确
		if login.User != "root" || login.Password != "admin" {
			c.JSON(http.StatusBadRequest, gin.H{"status": "304"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "200"})
	})
	r.Run(":8000")
}
