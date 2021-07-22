package main

import (
	"code.oldboyedu.com/go_5_blog/controller"
	"code.oldboyedu.com/go_5_blog/dao/db"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	dns := "root:123456@tcp(127.0.0.1:3306)/test?parseTime=true"
	err := db.Init(dns)
	if err != nil {
		panic(err)
	}
	// 加载静态文件
	router.Static("/static/", "./static")
	// 加载模板
	router.LoadHTMLGlob("views/*")
	router.GET("/", controller.IndexHandle)
	router.GET("/category/", controller.CategoryList)
	router.Run(":8000")
}
