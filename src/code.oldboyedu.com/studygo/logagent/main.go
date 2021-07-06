package main

import (
	"code.oldboyedu.com/studygo/logagent/kafka"
	"code.oldboyedu.com/studygo/logagent/taillog"
	"fmt"
	"time"
)

// LogAgent入口程序

func run() {
	// 1.读取日志
	for {
		select {
		case line := <-taillog.ReadChan():
			// 2.发送到kafka
			kafka.SendToKafka("web_log", line.Text)
		default:
			time.Sleep(time.Second)
		}
	}
}

func main() {
	// 1.初始化kafka连接
	err := kafka.Init([]string{"127.0.0.1:9092"})
	if err != nil {
		fmt.Printf("init kafka failed, err:%v\n", err)
		return
	}
	fmt.Println("init kafka success")
	// 2.打开日志文件准备收集日志
	err = taillog.Init("./my.log")
	if err != nil {
		fmt.Printf("init tailog failed, err:%v\n", err)
		return
	}
	fmt.Println("init tailog success")
	run()
}
