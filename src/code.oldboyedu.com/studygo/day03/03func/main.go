package main

import "fmt"

//函数

func f1() {
	fmt.Println("Hello 桂林！")
}

func f2(name string) {
	fmt.Println("Hello", name)

}

// 带参数和返回值的函数
func f3(x int, y int) int {
	sum := x + y
	return sum
}

// 参数类型简写
func f4(x, y int) int {
	return x + y
}

//可变参数
func f5(title string, y ...int) int {
	fmt.Println(y) // y是一个int类型的切片
	return 1
}

// 命名返回值
func f6(x, y int) (sum int) {
	sum = x + y //如果使用命名的返回值，那么在函数中可以直接使用返回值变量
	return      //如果使用命名的返回值，return后面可以省略返回值变量
}

// 多个返回值
func f7(x, y int) (sum, sub int) {
	sum = x + y
	sub = x - y
	return
}

func main() {
	f1()
	f2("荔浦")
	f3(100, 200)
	fmt.Println(f7(100, 200))

}
