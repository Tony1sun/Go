package main

import "fmt"

//结构体

type person struct {
	name   string
	age    int
	gender string
	hobby  []string
}

func main() {
	// 声明一个person类型的变量
	var p person
	// 通过字段赋值
	p.name = "今天"
	p.age = 18
	p.gender = "男"
	p.hobby = []string{"篮球", "足球", "双色球"}
	fmt.Println(p)
	// 访问变量p的字段
	fmt.Printf("%T\n", p)
	fmt.Println(p.name)

	var p2 person
	p2.name = "明天"
	p2.age = 30
	fmt.Printf("type:%T value:%v\n", p2, p2)

	// 匿名结构体
	var s struct {
		x string
		y int
	}
	s.x = "哈哈"
	s.y = 100
	fmt.Printf("type:%T value:%v\n", s, s)
}
