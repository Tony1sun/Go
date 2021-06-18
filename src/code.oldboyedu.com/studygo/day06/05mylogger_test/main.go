package main

import (
	"code.oldboyedu.com/studygo/day06/mylogger"
	"time"
)

// 测试写的日志库
func main() {
	log := mylogger.NewLog("info")
	for {
		log.Debug("Debug测试日志")
		log.Trace("trace测试日志")
		log.Info("Info测试日志")
		log.Warning("warning测试日志")
		log.Error("error测试日志")
		log.Fatal("fatal测试日志")
		time.Sleep(time.Second * 3)
	}

}
