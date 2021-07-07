package main

import (
	"code.oldboyedu.com/studygo/logagent1/conf"
	"code.oldboyedu.com/studygo/logagent1/kafka"
	"code.oldboyedu.com/studygo/logagent1/taillog"
	"fmt"
	"gopkg.in/ini.v1"
	"time"
)

var (
	cfg = new(conf.AppConf)
)

func run() {
	// 1.读取日志
	for {
		select {
		case line := <-taillog.ReadChan():
			// 发送到kafka
			kafka.SendToKafka(cfg.Topic, line.Text)
		default:
			time.Sleep(time.Second)
		}
	}
}

func main() {
	// 1. 加载配置文件
	err := ini.MapTo(cfg, "./conf/config.ini")
	if err != nil {
		fmt.Println("load ini failed, err:%v\n", err)
		return
	}
	// 2. 初始化kafka连接
	err = kafka.Init([]string{cfg.Address})
	if err != nil {
		fmt.Println("Init kafka failed, err:%v\n", err)
		return
	}
	fmt.Println("init kafka success")
	// 3.打开日志文件准备收集日志
	err = taillog.Init(cfg.FileName)
	if err != nil {
		fmt.Println("Init taillog failed, err:%v\n", err)
		return
	}
	fmt.Println("init taillog success")
	run()
}
