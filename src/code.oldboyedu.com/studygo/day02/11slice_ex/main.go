package main

import "fmt"

func main() {
	// 切片的练习题
	var a = make([]int, 5, 10) //创建切片， 长度为5， 容量为10
	fmt.Println(a)
	for i := 0; i < 10; i++ {
		a = append(a, i)
	}
	fmt.Println(a)
	fmt.Println(cap(a))
}
