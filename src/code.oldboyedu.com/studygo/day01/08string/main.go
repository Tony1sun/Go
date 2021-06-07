package main

import (
	"fmt"
	"strings"
)

//字符串
func main()  {
	path := "'D:\\Go\\src\\code.oldboyedu.com\\studygo\\day01'"
	fmt.Println(path)

	s := "i'm ok"
	fmt.Println(s)
	//多行字符串
	s2 := `
		今天
		明天
		后天
	`
	fmt.Println(s2)
	s3 := `D:\Go\src\code.oldboyedu.com\studygo\day01`
	fmt.Println(s3)

	//字符串相关操作
	fmt.Println(len(s3))

	//字符串拼接
	name := "今天"
	world := "明天"
	ss := name + world
	fmt.Println(ss)
	ss1 := fmt.Sprintf("%s%s", name, world)
	fmt.Println(ss1)

	//分割
	ret := strings.Split(s3, "\\")
	fmt.Println(ret)

	//包含
	fmt.Println(strings.Contains(ss, "今天"))
	fmt.Println(strings.Contains(ss, "明天"))

	//前缀
	fmt.Println(strings.HasPrefix(ss, "今"))
	//后缀
	fmt.Println(strings.HasSuffix(ss, "明"))

	s4 := "abcdeb"
	fmt.Println(strings.Index(s4, "c"))
	fmt.Println(strings.LastIndex(s4, "b"))

	//拼接
	fmt.Println(strings.Join(ret, "+"))
}