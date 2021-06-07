package main

import "fmt"

//单行注释

/*
多行注释
*/

//go语言函数外的语句必须以关键字开头
//var name = "哈哈"
//var age int
// age = 18

//const (
//	Num = 100
//)
//main函数是入口函数
//它没有参数也没有返回值
func main() {
	//函数内部定义的变量必须使用
	//var isOK = true
	//fmt.Print(isOK)
	//var age = 19
	//if age > 18 {
	//	fmt.Println("成年了")
	//} else if age > 7 {
	//	fmt.Println("娃仔卵")
	//} else {
	//	fmt.Println("小朋友")
	//}

	// for
	//for i:=0;i<10;i++ {
	//	fmt.Println(i)
	//}
	// for 2
	//for ;i<10;i++ {
	//	fmt.Println(i)
	//}
	//fmt.Println(i)	var i = 0

	// for 3
	//var j = 0
	//for j < 10 {
	//	fmt.Println(j)
	//	j++
	//}

	// for 4
	//for {
	//	fmt.Println("无限循环")
	//}

	// for 5
	//s := "hello"
	//fmt.Println(len(s))
	//for i, v := range s {
	//	fmt.Printf("%d %c\n", i, v)
	//}

	//哑元变量，不想用到的直接仍给它
	/*	for _, v := range s {
		fmt.Printf("%c\n", v)
	}*/

	// 九九乘法表
	//for i := 1;i < 10; i++ {
	//	for j := 1; j <= i; j++ {
	//		fmt.Printf("%d*%d=%d\t", j, i, i*j)
	//	}
	//	fmt.Println()
	//}

	//byte和rune
	//s1 := "Hello"
	//s2 := "今天明天后天"
	//for _,v := range s1 {
	//	fmt.Printf("%c %T\n", v, v)
	//}
	//for _,v := range s2 {
	//	fmt.Printf("%c %T\n", v, v)
	//}

	//八进制数
	var n1 = 0777

	//十六进制数
	var n2 = 0xff
	fmt.Print(n1, n2)
	fmt.Printf("%o\n", n1)
	fmt.Printf("% x\n", n2)

}
