package es

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"strings"
)

// 初始化ES， 准备接收kafka那边发来的数据

var (
	client *elastic.Client
)

// Init ...
func Init(address string) (err error) {
	if !strings.HasPrefix(address, "http://") {
		address = "http://" + address
	}
	client, err = elastic.NewClient(elastic.SetURL(address))
	if err != nil {
		return err
	}
	fmt.Println("connect to es success")
	return
}

// SendToEs 发送数据到Es
func SendToEs(indexStr string, data interface{}) (err error) {
	//

	// 链式操作
	put1, err := client.Index().
		Index(indexStr).
		BodyJson(data).
		Do(context.Background())
	if err != nil {
		return err
	}
	fmt.Printf("Indexed student %s to index %s, type %s\n", put1.Id, put1.Index, put1.Type)
	return err
}
