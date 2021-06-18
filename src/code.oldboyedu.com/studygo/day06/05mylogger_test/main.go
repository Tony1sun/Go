package main

import "code.oldboyedu.com/studygo/day06/mylogger"

// 测试写的日志库
func main() {
	log := mylogger.NewLog()
	for {
		log.Debug("Debug测试日志")
		log.Info("Info测试日志")
	}

}
