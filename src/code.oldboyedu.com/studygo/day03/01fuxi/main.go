package main

import "fmt"

func main() {
	//数组的初始化
	//var ages [30]int  //声明了一个变量ages,它是[30]int类型
	//ages = [30]int{1,2,3,4,5,6}  //初始化方式
	//fmt.Println(ages)
	//var ages2 = [...]int{1,2,3,4,5,6,7,8,9}  //声明一个变量的同时初始化
	//fmt.Println(ages2)
	//var ages3 = [...]int{1: 100, 10: 200}  //初始化的同时在指定位置插入指定的值
	//fmt.Println(ages3)
	//
	////数组的遍历-二维数组
	//var a1 = [...][2]int {//[[1 2] [3 4] [5 6]]
	//	[2]int{1, 2},
	//	[2]int{3, 4},
	//	[2]int{5, 6},
	//}
	//fmt.Println(a1)
	//多维数组只有最外层可以使用...

	//切片
	//var s1 []int // 声明时没有分配内存 == nil
	//s1 = []int{1, 2, 3}
	//fmt.Println(s1)
	//// make初始化 分配内存
	//s2 := make([]bool, 2, 4)
	//fmt.Println(s2)
	//s3 := make([]int, 0, 4)
	//fmt.Println(s3 == nil)

	//s1 := []int{1, 2, 3}
	//s2 := s1
	//fmt.Println(s2)
	//s2[1] = 200
	//fmt.Println(s2)
	//fmt.Println(s1)

	//拷贝
	//s1 := []int{1, 2, 3}
	//s2 := s1
	//var s3 = make([]int, 3, 3)
	//copy(s3, s1)
	//fmt.Println(s2)
	//s2[1] = 200
	//fmt.Println(s2)
	//fmt.Println(s1)
	//fmt.Println(s3)

	//指针
	//Go里面的指针只能读不能修改
	//addr := "深圳"
	//addrP := &addr
	//fmt.Println(addrP)
	//fmt.Printf("%T\n", addrP)
	//addrV := *addrP
	//fmt.Println(addrV)

	// map
	var m1 map[string]int
	fmt.Println(m1 == nil)
	m1 = make(map[string]int, 10)
	m1["今天"] = 66
	fmt.Println(m1)
	fmt.Println(m1["fadfs"]) //如果娶不到，就返回0
	score, ok := m1["今天"]
	if !ok {
		fmt.Println("不存在")
	} else {
		fmt.Println("分数是", score)
	}
	delete(m1, "lixiang") //删除的key不存在什么都不干
	delete(m1, "今天")
	fmt.Println(m1)
	fmt.Println(m1 == nil) //已经开辟了内存
}
