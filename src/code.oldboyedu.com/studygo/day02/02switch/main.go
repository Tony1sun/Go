package main

import "fmt"

//switch语句
func main() {
	//var n = 3
	//switch n :=3; n {
	//case 1:
	//	fmt.Println("大拇指")
	//case 2:
	//	fmt.Println("食指")
	//case 3:
	//	fmt.Println("种植")
	//case 4:
	//	fmt.Println("无名指")
	//case 5:
	//	fmt.Println("小指")
	//default:
	//	fmt.Println("无效的数据")
	//
	//}

	switch n := 7; n {
	case 1, 3, 5, 7, 9:
		fmt.Println("奇数")
	case 2, 4, 6, 8:
		fmt.Println("偶数")
	default:
		fmt.Println("n")

	}
}
