package taillog

import (
	"code.oldboyedu.com/studygo/logagent/kafka"
	"fmt"
	"github.com/hpcloud/tail"
)

// 专门往kafka写日志的模块

// 一个日志收集的任务
type TailTask struct {
	path     string
	topic    string
	instance *tail.Tail
}

func NewTailTask(path, topic string) (tailObj *TailTask) {
	tailObj = &TailTask{
		path:  path,
		topic: topic,
	}
	tailObj.init() // 根据路径去打开对应的日志
	return
}

func (t TailTask) init() {
	config := tail.Config{
		ReOpen:    true,                                 // 重新打开
		Follow:    true,                                 // 是否跟随
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2}, // 从文件的哪个地方开始读
		MustExist: false,                                // 文件不存在不报错
		Poll:      true,
	}
	var err error
	t.instance, err = tail.TailFile(t.path, config)
	if err != nil {
		fmt.Println("tail file failed, err:", err)
	}
	go t.run()
}

func (t *TailTask) run() {
	for {
		select {
		case line := <-t.instance.Lines: // 从tailObj通道中一行一行读取日志数据
			// 3.2发往kafka
			//kafka.SendToKafka(t.topic, line.Text)
			// 先把日志数据发到一个通道中
			kafka.SendToChan(t.topic, line.Text)
			// kafka 那个包中有单独的goroutine去取日志数据发到kafka
		}
	}
}
