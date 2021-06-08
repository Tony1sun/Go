package main

import "fmt"

func test1(name string) {
	fmt.Println("Hello,", name)
}

// 函数作为参数
func test(f func(string), name string) {
	f(name)
}

// 函数作为返回值
func test2() func(int, int) int {
	return func(x, y int) int {
		return x + y
	}
}

func low(f func()) {
	f()
}

// 闭包
func bi(f func(string), name string) func() {
	return func() {
		f(name)
	}
}

func main() {
	test(test1, "今天")
	ret := test2()
	fmt.Printf("%T\n", ret)
	sum := ret(10, 20)
	fmt.Println(sum)
	// 闭包示例
	fc := bi(test1, "今天")
	low(fc)
}
