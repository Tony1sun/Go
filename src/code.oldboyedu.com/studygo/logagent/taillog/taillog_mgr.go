package taillog

import "code.oldboyedu.com/studygo/logagent/etcd"

var tskMgr *tailLogMgr

type tailLogMgr struct {
	LogEntry []*etcd.LogEntry
	//tailMap map[string]*TailTask
}

func Init(logConf []*etcd.LogEntry) {
	tskMgr = &tailLogMgr{
		LogEntry: logConf, // 把当前的日志收集项配置保存
	}
	for _, logEntry := range logConf {
		// conf: *etcd.LogEntry
		// LogEntry.Path 要收集的日志路径
		NewTailTask(logEntry.Path, logEntry.Topic)
	}
}
