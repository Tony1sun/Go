package es

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"strings"
	"time"
)

// 初始化ES， 准备接收kafka那边发来的数据

type LogData struct {
	Topic string `json:"topic"`
	Data  string `json:"data"`
}

var (
	client *elastic.Client
	ch     chan *LogData
)

// Init ...
func Init(address string, chanSize, nums int) (err error) {
	if !strings.HasPrefix(address, "http://") {
		address = "http://" + address
	}
	client, err = elastic.NewClient(elastic.SetURL(address))
	if err != nil {
		return err
	}
	fmt.Println("connect to es success")
	ch = make(chan *LogData, chanSize)
	for i := 0; i < nums; i++ {
		go sendToEs()
	}
	return
}

// SendToEs 发送数据到Es
func SendToEsChan(msg *LogData) {
	//
	ch <- msg

}

func sendToEs() error {
	// 链式操作
	for {
		select {
		case msg := <-ch:
			put1, err := client.Index().
				Index(msg.Topic).
				BodyJson(msg).
				Do(context.Background())
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Printf("Indexed student %s to index %s, type %s\n", put1.Id, put1.Index, put1.Type)
		default:
			time.Sleep(time.Second)
		}
	}

}
