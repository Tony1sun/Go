package main

import "fmt"

// Go语言中函数的return 不是原子操作，在底层是分为两步来执行
// 第一步：返回值赋值
// defer
// 第二部：真正的RET返回
// 函数中如果存在 defer， 那么defer执行的时机实在第一步和第二部之间

func f1() int {
	x := 5
	defer func() {
		x++ //修改的是x，不是返回值
	}()
	return x
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5
}

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x
}

func main() {
	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
}
