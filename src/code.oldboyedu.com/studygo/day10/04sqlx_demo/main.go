package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// Go连接MySQL示例
var db *sqlx.DB // 是一个连接池对象

func initDB() (err error) {
	// 数据库信息
	dsn := "root:123456@tcp(127.0.0.1:3306)/sql_test"
	// 连数据库
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		return
	}

	db.SetMaxOpenConns(10) // 设置数据库连接池的最大连接数
	db.SetMaxIdleConns(10) // 设置最大空闲连接数
	return
}

type user struct {
	ID   int
	Name string
	Age  int
}

func sqqueryone(id int) {
	sqlStr1 := `select id, name, age from user where id=?`
	var u user
	db.Get(&u, sqlStr1, id)
	fmt.Printf("u:%#v\n", u)
}

func sqquerymore() {
	var userList []user
	sqlStr2 := `select id, name, age from user`
	err := db.Select(&userList, sqlStr2)
	if err != nil {
		fmt.Printf("err:%#v\n", err)
		return
	}
	fmt.Printf("userlist:%#v\n", userList)
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init DB failed, err:%v\n", err)
		return
	}
	sqqueryone(6)
	sqquerymore()
}
