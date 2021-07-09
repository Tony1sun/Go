package main

import (
	"code.oldboyedu.com/studygo/logagent/conf"
	"code.oldboyedu.com/studygo/logagent/etcd"
	"code.oldboyedu.com/studygo/logagent/kafka"
	"code.oldboyedu.com/studygo/logagent/taillog"
	"code.oldboyedu.com/studygo/logagent/utils"
	"fmt"
	"gopkg.in/ini.v1"
	"sync"
	"time"
)

// LogAgent入口程序
var (
	cfg = new(conf.AppConf)
)

func main() {
	// 0.加载配置文件
	err := ini.MapTo(cfg, "./conf/config.ini")
	if err != nil {
		fmt.Println("load ini failed, err:%v\n", err)
		return
	}
	// 1.初始化kafka连接
	err = kafka.Init([]string{cfg.KafkaConf.Address}, cfg.KafkaConf.ChanMaxSize)
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
	// 为了实现每个logagentagent都拉去自己的独有配置，所以要以自己的IP地址作为区分
	ipStr, err := utils.GetOutboundIP()
	if err != nil {
		panic(err)
	}
	etcdConfkey := fmt.Sprintf(cfg.EtcdConf.Key, ipStr)
	// 2.1 从etcd中获取日志收集项的配置
	logEntryConf, err := etcd.GetConf(etcdConfkey)
	if err != nil {
		fmt.Printf("etcd.GetCond failed, err:%v\n", err)
		return
	}
	fmt.Printf("get conf from etcd success, %v\n", logEntryConf)
	// 2.1 派一个哨兵监视日志收集项的变化

	for index, value := range logEntryConf {
		fmt.Printf("index:%v value:%v\n", index, value)
	}
	// 3.收集日志发往kafka
	taillog.Init(logEntryConf)           // 因为 NewConfChan 访问了tskMgr 的 NewConfChan
	newConfChan := taillog.NewConfChan() // 从taillog包中获取对外暴露的通道
	var wg sync.WaitGroup
	wg.Add(1)
	go etcd.WatchConf(etcdConfkey, newConfChan) // 哨兵发现最新的配置信息会通知上面的那个通道
	wg.Wait()
}
