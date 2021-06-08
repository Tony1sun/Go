package main

import "fmt"

// 结构体是值类型

type person struct {
	name, gender string
}

func f(x person) {
	x.gender = "女"
}

func f2(x *person) {
	x.gender = "女" // 语法糖，自动根据指针找对应的变量
}
func main() {
	var p person
	p.name = "兰博"
	p.gender = "男"
	f(p)
	fmt.Println(p.gender) // 男
	f2(&p)
	fmt.Println(p.gender) // 女
	// 结构体指针1
	var p2 = new(person)
	p2.name = "理想"
	p2.gender = "男"
	fmt.Printf("%T\n", p2)
	fmt.Printf("%p\n", p2)  // p2保持的值就是一个内存地址
	fmt.Printf("%p\n", &p2) // p2的内存地址
	// 2. 结构体指针2
	// 2.1 key-value初始化
	var p3 = person{
		name:   "元帅",
		gender: "男",
	}
	fmt.Printf("%#v\n", p3)
	// 2.2 使用值列表的形式初始化
	p4 := person{
		"小王子",
		"男",
	}
	fmt.Printf("%#v\n", p4)
}
