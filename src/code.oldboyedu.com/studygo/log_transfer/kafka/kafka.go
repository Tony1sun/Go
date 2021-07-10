package kafka

import (
	"code.oldboyedu.com/studygo/log_transfer/es"
	"fmt"
	"github.com/Shopify/sarama"
)

// 初始化kafka消费者 从kafka取数据
type LogData struct {
	data string `json:"data"`
}

// Init初始化client
func Init(addrs []string, topic string) (err error) {
	consumer, err := sarama.NewConsumer(addrs, nil)
	if err != nil {
		fmt.Printf("fail to start consumer, err:%v\n", err)
		return err
	}
	partitionList, err := consumer.Partitions(topic) // 根据topic取到所有的分区
	if err != nil {
		fmt.Printf("fail to get list of partition:err:%v\n", err)
		return err
	}
	fmt.Println("分区列表:", partitionList)
	for partition := range partitionList { // 遍历所有的分区
		// 针对每个分区创建一个对应的分区消费者
		pc, err := consumer.ConsumePartition(topic, int32(partition), sarama.OffsetNewest)
		if err != nil {
			fmt.Printf("fail to start sonsumer for partition %d,err:%v\n", partition, err)
			return err
		}
		//defer pc.AsyncClose()
		// 异步从每个分区消费信息
		go func(partitionConsumer sarama.PartitionConsumer) {
			for msg := range pc.Messages() {
				fmt.Printf("Partition:%d Offset:%d Key:%v Value:%v\n", msg.Partition, msg.Offset, msg.Key, string(msg.Value))
				// 直接发给ES
				ld := map[string]interface{}{
					"data": string(msg.Value),
				}
				if err != nil {
					fmt.Printf("unmarshal failed, err:%v\n", err)
					continue
				}
				es.SendToEs(topic, ld)
			}
		}(pc)
	}
	return err
}
