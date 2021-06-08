package main

import "fmt"

func main() {
	//fmt.Print("今天")
	//fmt.Print("今天")
	//fmt.Println("------")
	//fmt.Println("今天")
	//fmt.Printf("%")
	// Printf("格式化字符串", 值)
	// %T：查看类型
	// %d：十进制数
	// %b：二进制数
	// %o：八进制数
	// %x：十六进制
	// %c: 字符
	// %s:字符串
	// %p:指针
	// %v：值
	// %f:浮点数

	//var m1 = make(map[string]int, 1)
	//m1["今天"] = 100
	//fmt.Printf("%v\n", m1)
	//fmt.Printf("%#v\n", m1)
	//
	//fmt.Printf("%v\n", 100)
	//// 整数 -> 字符
	//fmt.Printf("%q\n", 65)
	////浮点数和复数
	//fmt.Printf("%b\n", 3.141592953)
	//// 字符串
	//fmt.Printf("%q\n", "今天明天后天")
	//fmt.Printf("%1.3s\n", "今天明天后天")

	//获取用户输入
	//var s string
	//fmt.Scan(&s)
	//fmt.Println("输入的内容是：", s)

	var (
		name  string
		age   int
		class string
	)
	//fmt.Scanf("%s %d %s\n", &name, &age, &class)
	//fmt.Println(name, age, class)

	fmt.Scanln(&name, &age, &class)
	fmt.Println(name, age, class)

}
