package main

import "fmt"

//函数的定义
func sum(x int, y int) (ret int) {
	return x + y
}

//没有返回值
func f1(x int, y int) {
	fmt.Println(x + y)
}

//没有参数没有返回值
func f2() {
	fmt.Println("f2")
}

//没有参数但有返回值
func f3() int {
	ret := 3
	return ret
}

//返回值可以命名也可以不命名

//命名的返回值就相当于在函数中声明一个变量
func f4(x int, y int) (ret int) {
	ret = x + y
	return
}

//多个返回值
func f5() (int, string) {
	return 1, "今天"
}

//参数的类型简写
//当参数中连续多个参数的类型一致时，我们可以只写最后一个参数的类型
func f6(x, y, z int, m, n string, i, j bool) int {
	return x + y
}

//可变长参数
//可变长参数必须放在函数参数的最后
func f7(x string, y ...int) {
	fmt.Println(x)
	fmt.Println(y)
}

//Go语言中函数没有默认参数这个概念

func main() {
	a1 := sum(1, 2)
	fmt.Println(a1)
	_, n := f5()
	fmt.Println(n)

	f7("下雨了", 1, 2, 3, 4, 5, 6)
}
