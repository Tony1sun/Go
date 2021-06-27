package main

import (
	"fmt"
	"sync"
)

var a []int
var b chan int //需要指定通道中元素的类型
var wg sync.WaitGroup

func noBufChannel() {
	b = make(chan int) // 通道的初始化 ps:通道需要初始化才能使用
	wg.Add(1)
	go func() {
		defer wg.Done()
		x := <-b
		fmt.Println("后台goroutine从通道b中取到了", x)
	}()
	b <- 10
	fmt.Println("10发送到通道b中了...")
	wg.Wait()
}

func BufChannel() {
	b = make(chan int, 10) // 通道的初始化 ps:通道需要初始化才能使用
	b <- 10
	fmt.Println("10发送到通道b中了...")
	b <- 20
	fmt.Println("20发送到通道b中了...")
	x := <-b
	fmt.Println("从通道b中取到了", x)
	close(b)

}

func main() {
	BufChannel()
	//noBufChannel()
}
