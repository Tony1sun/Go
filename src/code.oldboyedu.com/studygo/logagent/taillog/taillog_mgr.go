package taillog

import "code.oldboyedu.com/studygo/logagent/etcd"

var tskMgr *tailLogMgr

type tailLogMgr struct {
	LogEntry []*etcd.LogEntry
	tskMap map[string]*TailTask
	newConfChan chan []*etcd.LogEntry
}

func Init(logConf []*etcd.LogEntry) {
	tskMgr = &tailLogMgr{
		LogEntry: logConf, // 把当前的日志收集项配置保存
		tskMap: make(map[string]*TailTask, 32),
		newConfChan: make(chan []*etcd.LogEntry), // 无缓冲区通道
	}
	for _, logEntry := range logConf {
		// conf: *etcd.LogEntry
		// LogEntry.Path 要收集的日志路径
		NewTailTask(logEntry.Path, logEntry.Topic)
	}
}

// 监听自己的NewConfChan，有了新配置过来之后就做对应处理
// 1. 配置新增
// 2. 配置删除
// 3. 配置变更
func (t *tailLogMgr) run() {
	for {
		select {
		case :
			
		}
	}
}
