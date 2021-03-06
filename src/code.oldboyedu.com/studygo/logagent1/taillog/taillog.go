package taillog

import (
	"fmt"
	"github.com/hpcloud/tail"
)

var (
	tailObj *tail.Tail
	LogChan chan string
)

func Init(filename string) (err error) {
	config := tail.Config{
		ReOpen:    true,                                 // 重新打开
		Follow:    true,                                 // 是否跟随
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2}, // 从文件的哪个地方开始
		MustExist: false,                                // 文件不存在不报错
		Poll:      true,
	}
	tailObj, err = tail.TailFile(filename, config)
	if err != nil {
		fmt.Println("tail file failed, err:", err)
		return
	}
	return
}

func ReadChan() <-chan *tail.Line {
	return tailObj.Lines
}
