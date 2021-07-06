package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
)

var (
	client sarama.SyncProducer // 声明一个全局的连接kafka的生产者client
)

// Init初始化client
func Init(addrs []string) (err error) {
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
	return
}

func SendToKafka(topic, data string) {
	//构造一个消息
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(data),
	}
	// 发送到kafka
	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		fmt.Println("send msg failed , err:", err)
		return
	}
	fmt.Printf("pid :%v offset :%v\n", pid, offset)
}