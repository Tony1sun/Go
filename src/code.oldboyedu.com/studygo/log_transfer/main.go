package main

import (
	"code.oldboyedu.com/studygo/log_transfer/conf"
	"code.oldboyedu.com/studygo/log_transfer/es"
	"code.oldboyedu.com/studygo/log_transfer/kafka"
	"fmt"
	"gopkg.in/ini.v1"
)

// log transfer
// 将日志数据从kafka取出来发往ES

func main() {
	// 0. 加载配置文件
	var (
		cfg = new(conf.LogTransferCfg)
	)
	err := ini.MapTo(cfg, "./conf/cfg.ini")
	if err != nil {
		fmt.Println("init config failed, err:%v\n", err)
		return
	}
	fmt.Printf("cfg:%v\n", cfg)
	// 1.初始化ES
	// 1.1初始化一个ES连接的client
	err = es.Init(cfg.EsCfg.Address)
	if err != nil {
		fmt.Printf("init ES client failed, err:%v\n", err)
		return
	}
	fmt.Println("init es success")
	// 2. 初始化kafka
	// 2.1 连接kafka， 创建分区的消费者
	// 2.2每个分区的消费者分别取出数据通过sendToEs将数据发往ES
	err = kafka.Init([]string{cfg.KafkaCfg.Address}, cfg.KafkaCfg.Topic)
	if err != nil {
		fmt.Printf("init kafka consumer failed, err:%v\n", err)
		return
	}
	select {}
	// 1. 从kafka取日志数据

	// 2. 发往ES
}
