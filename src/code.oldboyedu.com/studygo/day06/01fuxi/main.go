package main

import "fmt"

func main() {
	var a interface{} // 定义一个空接口变量
	a = 100
	//v1, ok := a.(int8)
	//if ok {
	//	fmt.Println("猜对了,a是int8", v1)
	//	return
	//}else {
	//	fmt.Println("猜错了!")
	//}

	// switch
	switch v2 := a.(type) {
	case int8:
		fmt.Println("int8", v2)
	case int16:
		fmt.Println("int16", v2)
	case int:
		fmt.Println("int", v2)
	case string:
		fmt.Println("string", v2)
	default:
		fmt.Println("滚~")
	}

}
