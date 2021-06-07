package main

import "fmt"

// 闭包
func f1(f func()) {
	fmt.Println("is f1")
	f()
}

func f2(x, y int) {
	fmt.Println("is f2")
	fmt.Println(x + y)
}

//定义一个函数对f2进行包装
//func f3() func(int, int) {
//	tmp := func(){
//
//	}
//}

func f3(f func(int, int), x, y int) func() {
	tmp := func() {
		f(x, y)
	}
	return tmp
}
func main() {
	ret := f3(f2, 100, 200)
	f1(ret)
}
