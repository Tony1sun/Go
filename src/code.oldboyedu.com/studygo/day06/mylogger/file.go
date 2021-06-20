package mylogger

import (
	"fmt"
	"time"
)

// 往文件里面写日志相关代码

type FileLogger struct {
	Level       LogLevel
	filePath    string // 日志文件保存的路径
	fileName    string // 日志文件名字
	maxFileSize int64
}

// NewFileLogger
func NewFileLogger(leverStr, fp, fn string, maxSize int64) *FileLogger {
	loglevel, err := parseLogLevel(leverStr)
	if err != nil {
		panic(err)
	}
	return &FileLogger{
		Level:       loglevel,
		filePath:    fp,
		fileName:    fn,
		maxFileSize: maxSize,
	}
}

func (l ConsoleLogger) enable(logLevel LogLevel) bool {
	return l.Level <= logLevel

}

func log(lv LogLevel, format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...)
	now := time.Now()
	funcName, fileName, lineNo := getInfo(3)
	fmt.Printf("[%s] [%s] [%s:%s:%d] %s\n", now.Format("2006-01-02 15:04:05"), getLogString(lv), fileName, funcName, lineNo, msg)

}

func (l ConsoleLogger) Debug(format string, a ...interface{}) {
	if l.enable(DEBUG) {
		log(DEBUG, format, a...)
	}
}

func (l ConsoleLogger) Trace(format string, a ...interface{}) {
	if l.enable(TRACE) {
		log(TRACE, format, a...)
	}
}

func (l ConsoleLogger) Info(format string, a ...interface{}) {
	if l.enable(INFO) {
		log(INFO, format, a...)
	}
}

func (l ConsoleLogger) Warning(format string, a ...interface{}) {
	if l.enable(WARNING) {
		log(WARNING, format, a...)
	}
}

func (l ConsoleLogger) Error(format string, a ...interface{}) {
	if l.enable(ERROR) {
		log(ERROR, format, a...)
	}
}

func (l ConsoleLogger) Fatal(format string, a ...interface{}) {
	if l.enable(FATAL) {
		log(FATAL, format, a...)
	}
}
