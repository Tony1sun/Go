package main

import "fmt"

func main()  {
	s := "Hello哈哈"
	//len()求的是byte字节的数量
	n := len(s)
	fmt.Println(n)


	//for _, c := range s {
	//	fmt.Printf("%c\n", c)
	//}

	//字符串修改
	s2 := "白萝卜" //
	s3 := []rune(s2)  //把字符串强制转换成了一个rune切片
	s3[0] = '红'
	fmt.Println(string(s3))

	c1 := "红"
	c2 := '红'  //rune(int32)
	fmt.Printf("c1:%T c2:%T\n", c1, c2)

	c3 := "H"
	c4 := byte('H')  //rune(int32)
	fmt.Printf("c3:%T c4:%T\n", c3, c4)
	fmt.Printf("%d\n", c4)

	//类型转换
	n1 := 10 //int
	var f float64
	f = float64(n1)
	fmt.Println(f)
	fmt.Printf("%T\n", f)
}
