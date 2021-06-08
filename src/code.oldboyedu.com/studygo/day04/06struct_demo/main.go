package main

import "fmt"

// 结构体占用一块连续的内存空间

type x struct {
	a int8
	b int8
	c string
}

func main() {
	m := x{
		a: int8(10),
		b: int8(20),
		c: "哈哈", // 有一个类型与其他类型不同时会内存对齐
	}
	fmt.Printf("%p\n", &(m.a))
	fmt.Printf("%p\n", &(m.b))
	fmt.Printf("%p\n", &(m.c))
}
