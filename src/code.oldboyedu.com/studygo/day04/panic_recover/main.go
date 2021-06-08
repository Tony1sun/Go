package main

import "fmt"

//panic 和 recover

func f1() {
	defer func() {
		err := recover() // 收集当前的错误
		fmt.Println("继续执行")
		fmt.Println(err)
	}()
	panic("致命错误")
	fmt.Println("f1")
}

func f2() {
	fmt.Println("f2")
}

func main() {
	f1()
	f2()
}
