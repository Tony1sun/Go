package main

import (
	"fmt"
	"io"
	"os"
)

// 文件操作
func f1() {
	var fileObj *os.File
	var err error
	fileObj, err = os.Open("./main.go")
	if err != nil {
		fmt.Printf("open file failed, err:%v", fileObj)
		return
	}
	defer fileObj.Close()
}

// Copyfile 拷贝文件函数
func CopyFile(dstName, srcName string) (written int64, err error) {
	// 以读方式打开源文件
	src, err := os.Open(srcName)
	if err != nil {
		fmt.Printf("open %s failed, err:%v.\n", srcName, err)
		return
	}
	defer src.Close()
	// 以写|创建的方式打开目标文件
	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Printf("open %s failed, err:%v.\n", dstName, err)
		return
	}
	defer dst.Close()
	return io.Copy(dst, src) // 调用io.Copy()拷贝内容
}

func f2() {
	// 打开要操作的文件
	fileObj, err := os.OpenFile("./sb.txt", os.O_RDWR, 0644)
	if err != nil {
		fmt.Printf("read from file failed, err:%v", err)
		return
	}
	defer fileObj.Close()
	// 因为没有办法直接在文件中间插入内容，所以要借助一个临时文件
	tmpFile, err := os.OpenFile("./sb.tmp", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("create tmp file failed, err:%v\n", err)
		return
	}
	defer tmpFile.Close()
	// 读取文件写入临时文件
	var ret [1]byte
	n, err := fileObj.Read(ret[:])
	if err != nil {
		fmt.Printf("read from file failed, err:%v", err)
		return
	}
	// 写入临时文件
	tmpFile.Write(ret[:n])
	// 再写入要插入的内容
	var s []byte
	s = []byte{'c'}
	tmpFile.Write(s)
	// 紧接着把原文件后续的内容写入临时文件
	var x [1024]byte
	for {
		n, err := fileObj.Read(x[:])
		if err == io.EOF {
			tmpFile.Write(x[:n])
			break
		}
		if err != nil {
			fmt.Printf("read from file failed, err:%v", err)
			return
		}
		tmpFile.Write(x[:n])
	}
	// 源文件后续也写入了临时文件
	fileObj.Close()
	tmpFile.Close()
	os.Rename("./sb.tmp", "./sb.txt")

	//n, err := fileObj.Read(ret[:])
	//if err != nil {
	//	fmt.Printf("read from file failed, err:%v", err)
	//	return
	//}
	//fmt.Println(string(ret[:n]))
}

func main() {
	//f1()
	//srcFile := "D:/realty测试数据/身份证/432432xixi.jpg"
	//dstFile := "D:/realty测试数据/身份证/微信图片_20210107102630.jpg"
	//srcFile := "D:/realty测试数据/身份证/test1.txt"  // 复制之后的文件
	//dstFile := "D:/realty测试数据/身份证/test.txt"  // 已存在的文件
	//_, err := CopyFile(srcFile, dstFile)
	//if err != nil {
	//	fmt.Println("copy file faile, err:", err)
	//	return
	//}
	//fmt.Println("copy done!")
	f2()
}
