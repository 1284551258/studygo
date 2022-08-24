package etcd

import (
	"context"
	"encoding/json"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
)

var cli *clientv3.Client

type LogEntry struct {
	Path  string `json:"path"`
	Topic string `json:"topic"`
}

func Init(addr string, timeout time.Duration) (err error) {
	cli, err = clientv3.New(clientv3.Config{
		Endpoints:   []string{addr},
		DialTimeout: timeout,
	})
	if err != nil {
		fmt.Printf("connect to etcd failed,err:%v", err)
		return
	}
	return
}
func GetConf(key string) (logEntryConf []*LogEntry, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.Get(ctx, key)
	cancel()
	if err != nil {
		fmt.Printf("get from etcd failed,err:%v\n", err)
		return
	}
	for _, ev := range resp.Kvs {
		err := json.Unmarshal(ev.Value, &logEntryConf)
		if err != nil {
			fmt.Printf("Unmarshal to logEntryConf failed,err:%v\n", err)
			return nil, err
		}
	}
	return
}

func WatchConf(key string, confChan chan<- []*LogEntry) {
	watch := cli.Watch(context.Background(), key)
	for response := range watch {
		for _, event := range response.Events {
			fmt.Printf("event:%v,key:%v,value:%v\n", event.Type, string(event.Kv.Key),
				string(event.Kv.Value))
			var newConf []*LogEntry
			if event.Type != clientv3.EventTypeDelete {
				err := json.Unmarshal(event.Kv.Value, &newConf)
				if err != nil {
					fmt.Printf("Unmarshal newConf failed,err:%v", err)
					continue
				}
			}
			confChan <- newConf
			fmt.Println("发送成功")
		}
	}

}

//./my1.log  web_log
//./my2.log  redis_log
