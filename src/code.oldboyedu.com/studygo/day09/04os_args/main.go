package main

import (
	"flag"
	"fmt"
)

// flag 获取命令行参数

func main() {
	// 创建一个标志位参数
	//name := flag.String("name", "今天", "名字")
	//age := flag.Int("age", 9000, "请输入真是年龄")
	//married := flag.Bool("married", false, "是否结婚")
	//cTime := flag.Duration("ct", time.Second, "结婚多久")

	var name string
	flag.StringVar(&name, "name", "今天", "名字")
	// 使用flag
	flag.Parse()
	fmt.Println(name)
	//fmt.Println(*name)
	//fmt.Println(*age)
	//fmt.Println(*married)
	//fmt.Printf("%T\n", *cTime)
	flag.Args()
	flag.NArg()
	flag.NFlag()
}
