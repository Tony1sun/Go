package main

import "fmt"

//defer多用于函数结束之前释放资源(文件句柄、数据库链接、socket连接)
func defreDemo() {
	fmt.Println("start")
	defer fmt.Println("嘿嘿嘿") // defer把它后面的语句延迟到函数即将返回的时候再执行
	defer fmt.Println("呵呵呵") // 一个函数中可以有多个defer语句
	defer fmt.Println("哈哈哈") // 多个derfer语句按照先进后出的顺序延迟执行
	fmt.Println("end")
}

func main() {
	defreDemo()
}
