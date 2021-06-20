package main

import (
	"code.oldboyedu.com/studygo/day06/mylogger"
	"time"
)

// 测试写的日志库
func main() {
	log := mylogger.NewLog("error")
	for {
		log.Debug("Debug测试日志")
		log.Trace("trace测试日志")
		log.Info("Info测试日志")
		log.Warning("warning测试日志")
		id := 001
		name := "tony"
		log.Error("error测试日志 id:%d,name:%s", id, name)
		log.Fatal("fatal测试日志")
		time.Sleep(time.Second * 3)
	}

}
