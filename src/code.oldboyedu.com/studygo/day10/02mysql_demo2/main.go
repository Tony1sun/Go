package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

// Go连接MySQL示例
var db *sql.DB // 是一个连接池对象

func initDB() (err error) {
	// 数据库信息
	dsn := "root:123456@tcp(127.0.0.1:3306)/goday10"
	// 连数据库
	db, err = sql.Open("mysql", dsn) // 不会校验用户名和密码是否正确
	if err != nil {
		return
	}
	err = db.Ping() // 尝试连接数据库
	if err != nil {
		return
	}
	return
}

func query() {

}

func insert() {

}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init DB failed, err:%v\n", err)
	}
	fmt.Println("连接数据库成功!")
}
