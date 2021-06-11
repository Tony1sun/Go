package main

import "fmt"

// 结构体嵌套

type address struct {
	province string
	city     string
}

type workPlace struct {
	province string
	city     string
}
type person struct {
	name    string
	age     int
	address // 嵌套匿名结构体
	workPlace
}

type company struct {
	name string
	address
}

func main() {
	p1 := person{
		name: "今天",
		age:  18,
		address: address{
			province: "广西",
			city:     "桂林",
		},
	}
	//fmt.Println(p1)
	//fmt.Println(p1.name, p1.address.city)
	//fmt.Println(p1.city)
	fmt.Println(p1.address.city)
	fmt.Println(p1.workPlace.city)
}
