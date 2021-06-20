package main

import (
	"fmt"
	"os"
)

// 文件对象的类型
// 获取文件对象的详细信息

func main() {
	fileObj, err := os.Open("./main.go")
	if err != nil {
		fmt.Printf("open file failed, err:%v\n", err)
		return
	}
	fmt.Printf("%T\n", fileObj)
}
