package taillog

import (
	"code.oldboyedu.com/studygo/logagent/etcd"
	"fmt"
	"time"
)

var tskMgr *tailLogMgr

type tailLogMgr struct {
	LogEntry    []*etcd.LogEntry
	tskMap      map[string]*TailTask
	newConfChan chan []*etcd.LogEntry
}

func Init(logConf []*etcd.LogEntry) {
	tskMgr = &tailLogMgr{
		LogEntry:    logConf, // 把当前的日志收集项配置保存
		tskMap:      make(map[string]*TailTask, 32),
		newConfChan: make(chan []*etcd.LogEntry), // 无缓冲区通道
	}
	for _, logEntry := range logConf {
		// conf: *etcd.LogEntry
		// LogEntry.Path 要收集的日志路径
		// 初始化的时候起了多少个tailtask ，都要记下来，为了后续判断方便
		tailObj := NewTailTask(logEntry.Path, logEntry.Topic)
		mk := fmt.Sprintf("%s_%s", logEntry.Path, logEntry.Topic)
		tskMgr.tskMap[mk] = tailObj
	}
	go tskMgr.run()
}

// 监听自己的NewConfChan，有了新配置过来之后就做对应处理

func (t *tailLogMgr) run() {
	for {
		select {
		case newConf := <-t.newConfChan:
			// 1. 配置新增
			for _, conf := range newConf {
				mk := fmt.Sprintf("%s_%s", conf.Path, conf.Topic)
				_, ok := t.tskMap[mk]
				if ok {
					// 原来就有，不需要操作
					continue
				} else {
					// 新增的
					tailObj := NewTailTask(conf.Path, conf.Topic)
					t.tskMap[mk] = tailObj
				}
			}
			// 找出原来t.logEntry有，但是newConf中没有的，要删掉
			for _, c1 := range t.LogEntry { // 从原来的配置中一次拿出配置项
				isDelete := true
				for _, c2 := range newConf { //去新的配置中逐一进行比较
					if c2.Topic == c1.Topic && c2.Path == c1.Path {
						isDelete = false
						continue
					}
				}
				if isDelete {
					// 把c1对用的这个tailObj给停掉
					mk := fmt.Sprintf("%s_%s", c1.Path, c1.Topic)
					t.tskMap[mk].cancelFunc()
				}
			}
			// 2. 配置删除
			// 3. 配置变更
			fmt.Println("新的配置来了", newConf)
		default:
			time.Sleep(time.Second)
		}
	}
}

// 向外部暴露tskMgr的newConfChan
func NewConfChan() chan<- []*etcd.LogEntry {
	return tskMgr.newConfChan
}
