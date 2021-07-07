package main

import (
	"code.oldboyedu.com/studygo/logagent/conf"
	"code.oldboyedu.com/studygo/logagent/etcd"
	"code.oldboyedu.com/studygo/logagent/kafka"
	"fmt"
	"gopkg.in/ini.v1"
	"time"
)

// LogAgent入口程序
var (
	cfg = new(conf.AppConf)
)

//func run() {
//	// 1.读取日志
//	for {
//		select {
//		case line := <- taillog.ReadChan():
//			// 2.发送到kafka
//			kafka.SendToKafka(cfg.KafkaConf.Topic, line.Text)
//		default:
//			time.Sleep(time.Second)
//		}
//	}
//}

func main() {
	// 0.加载配置文件
	err := ini.MapTo(cfg, "./conf/config.ini")
	if err != nil {
		fmt.Println("load ini failed, err:%v\n", err)
		return
	}
	// 1.初始化kafka连接
	err = kafka.Init([]string{cfg.KafkaConf.Address})
	if err != nil {
		fmt.Printf("init kafka failed, err:%v\n", err)
		return
	}
	fmt.Println("init kafka success")
	// 2.初始化etcd
	err = etcd.Init(cfg.EtcdConf.Address, time.Duration(cfg.EtcdConf.Timeout)*time.Second)
	if err != nil {
		fmt.Printf("init etcd failed, err:%v\n", err)
		return
	}
	fmt.Println("init etcd success")
	// 2.打开日志文件准备收集日志
	//err = taillog.Init(cfg.TaillogConf.FileName)
	//if err != nil {
	//	fmt.Printf("init tailog failed, err:%v\n", err)
	//	return
	//}
	//fmt.Println("init tailog success")
	//run()
}
