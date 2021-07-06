package main

import (
	"code.oldboyedu.com/studygo/logagent/kafka"
	"code.oldboyedu.com/studygo/logagent/taillog"
	"fmt"
	"gopkg.in/ini.v1"
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
	// 0.加载配置文件
	cfg, err := ini.Load("./conf/config.ini")
	fmt.Println(cfg.Section("kafka").Key("address"))
	fmt.Println(cfg.Section("kafka").Key("topic"))
	fmt.Println(cfg.Section("taillog").Key("path"))
	// 1.初始化kafka连接
	err = kafka.Init([]string{"127.0.0.1:9092"})
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
