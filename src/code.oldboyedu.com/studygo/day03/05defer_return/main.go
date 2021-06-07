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
	return x // 返回值赋值  2. defer  3.真正的RET指令
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

func f4() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5 // 返回值  = x = 5
}

func f5() (x int) {
	defer func(x int) int {
		x++ // 改变的是函数中的x的副本
		return x
	}(x)
	return 5

}
func main() {
	fmt.Println(f1()) // 5
	fmt.Println(f2()) // 6
	fmt.Println(f3()) // 5
	fmt.Println(f4()) // 5
	fmt.Println(f5()) // 5
}
