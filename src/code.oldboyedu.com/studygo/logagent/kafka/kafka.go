package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
	"time"
)

type logData struct {
	topic string
	data  string
}

var (
	client      sarama.SyncProducer // 声明一个全局的连接kafka的生产者client
	logDataChan chan *logData
)

// Init初始化client
func Init(addrs []string, maxSize int) (err error) {
	config := sarama.NewConfig()

	// tailf 包使用
	config.Producer.RequiredAcks = sarama.WaitForAll          // 发送完数据需要 leader和follow都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 新选出一个 partition;s随机；3种  哈希
	//config.Producer.Partitioner = sarama.NewRoundRobinPartitioner // 新选出一个 partition;轮训
	config.Producer.Return.Successes = true // 成功交付的消息将在success_channel返回

	//连接kafka
	client, err = sarama.NewSyncProducer(addrs, config)
	if err != nil {
		fmt.Println("producer closed err:", err)
		return
	}
	// 初始化loadDataChan
	logDataChan = make(chan *logData, maxSize)
	// 开启后台的goroutine从通道中取数据发往kafka
	go sendToKafka()
	return
}

// 给外部暴露的一个函数，该函数只把日志数据发送到一个内部的channel中
func SendToChan(topic, data string) {
	msg := &logData{
		topic: topic,
		data:  data,
	}
	logDataChan <- msg
}

// 真正往kafka发送日志的函数
func sendToKafka() {
	for {
		select {
		case ld := <-logDataChan:
			//构造一个消息
			msg := &sarama.ProducerMessage{
				Topic: ld.topic,
				Value: sarama.StringEncoder(ld.data),
			}
			// 发送到kafka
			pid, offset, err := client.SendMessage(msg)
			if err != nil {
				fmt.Println("send msg failed , err:", err)
				return
			}
			fmt.Printf("pid :%v offset :%v\n", pid, offset)
		default:
			time.Sleep(time.Microsecond * 50)
		}
	}
}
