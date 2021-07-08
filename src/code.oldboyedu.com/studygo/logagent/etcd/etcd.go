package etcd

import (
	"context"
	"encoding/json"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
)

var (
	cli *clientv3.Client
)

// 需要收集的日志配置信息
type LogEntry struct {
	Path  string `json:"path"`
	Topic string `json:"topic"`
}

// 初始化ETCD
func Init(addr string, timeout time.Duration) (err error) {
	cli, err = clientv3.New(clientv3.Config{
		Endpoints:   []string{addr},
		DialTimeout: timeout,
	})
	if err != nil {
		fmt.Printf("connect to etcd failed, err:%v\n", err)
		return
	}
	return
}

// 从etcd中根据key获取配置项
func GetConf(key string) (LogEntryConf []*LogEntry, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.Get(ctx, key)
	cancel()
	if err != nil {
		fmt.Printf("get from etcd failed, err:%v\n", err)
		return
	}
	for _, ev := range resp.Kvs {
		err = json.Unmarshal(ev.Value, &LogEntryConf)
		if err != nil {
			fmt.Printf("unmarshal etcd value failed, err:%v\n", err)
			return
		}
	}
	return
}

func WatchConf(key string, newConfch chan<- []*LogEntry) {
	ch := cli.Watch(context.Background(), key)
	// 从通道尝试取值
	for wresp := range ch {
		for _, evt := range wresp.Events {
			fmt.Printf("Type:%v key:%v value:%v\n", evt.Type, string(evt.Kv.Key), string(evt.Kv.Value))
			// 通知 taillog.tskMgr
			// 先判断操作的类型
			var newConf []*LogEntry
			//if evt.Type == clientv3.EventTypeDelete {
			//	// 如果是删除操作

			//}
			err := json.Unmarshal(evt.Kv.Value, &newConf)
			if err != nil {
				fmt.Printf("unmarshal failed, err:%v\n", err)
				continue
			}
			fmt.Printf("get new conf:%v\n", newConf)
			newConfch <- newConf
		}
	}
}
