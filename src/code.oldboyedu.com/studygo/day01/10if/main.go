package main

import "fmt"

// if条件判断

func main() {
	//age := 19
	//if age > 18 { // 如果 age > 18就执行这个 {}中的代码
	//	fmt.Println("澳门首家线上赌场开业啦!")
	//} else { // 否则就执行这个 {}中的代码
	//	fmt.Println("年龄未满18岁！")
	//}

	//多个判断条件
	//if age > 35 {
	//	fmt.Println("人道中年")
	//} else if age > 18 {
	//	fmt.Println("靓仔")
	//} else {
	//	fmt.Println("小朋友，好好学习！")
	//}

	// 作用域
	// age变量此时只在if条件判断语句中生效
	if age := 19; age > 18 { // 如果 age > 18就执行这个 {}中的代码
		fmt.Println("澳门首家线上赌场开业啦!")
	} else { // 否则就执行这个 {}中的代码
		fmt.Println("年龄未满18岁！")
	}
	//fmt.Println(age)  //在这里无法找到 age
}
