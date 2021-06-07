package main

import "fmt"
//Go语言中推荐使用驼峰式命名
var student_name string
var studentName string
var StudentName string

//声明变量
//var name string
//var age int
//var isok bool
var (
	name string
	age  int
	isok bool
)

func main() {
	name = "理想"
	age = 18
	isok = true
	//Go语言中变量声明必须使用,不然无法编译
	fmt.Print(isok)              //在终端中输出药打印的内容
	fmt.Printf("name:%s\n", name) //%s占位符 使用name这个变量的值去替换占位符
	fmt.Println(age)     //打印完指定的内容之后会在后面加一个换行符

	//声明变量同事赋值
	var s1 string = "whb"
	fmt.Println(s1)

	//类型推导(根据值判断该变量是什么类型)
	var s2 = "20"
	fmt.Println(s2)

	//简短变量声明
	s3 := "哈哈哈"
	fmt.Println(s3)

}
