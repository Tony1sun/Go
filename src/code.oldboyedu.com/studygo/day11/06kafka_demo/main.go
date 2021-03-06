package main

import (
	"fmt"
	"github.com/Shopify/sarama"
)

// kafka demo

//基于sarama第三方库开发的kafka client
func main() {
	config := sarama.NewConfig()

	//tailf 包使用
	config.Producer.RequiredAcks = sarama.WaitForAll          // 发送完数据需要 leader和follow都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 新选出一个 partition;s随机；3种  哈希
	//config.Producer.Partitioner = sarama.NewRoundRobinPartitioner // 新选出一个 partition;轮训
	config.Producer.Return.Successes = true // 成功交付的消息将在success_channel返回
	//构造一个消息
	msg := &sarama.ProducerMessage{
		Topic: "web_log",
		Value: sarama.StringEncoder("this is a test log"),
	}
	//连接kafka
	client, err := sarama.NewSyncProducer([]string{"127.0.0.1:9092"}, config)
	if err != nil {
		fmt.Println("producer closed err:", err)
		return
	}
	fmt.Println("连接kafka成功!")
	defer client.Close()
	//发送消息
	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		fmt.Println("send msg failed , err:", err)
		return
	}
	fmt.Printf("pid :%v offset :%v\n", pid, offset)
	//pid :0 offset :1
}
