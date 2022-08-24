package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
)

func main() {
	//初始化client
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"11.2.2.128:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Printf("connect to etcd failed,err:%v", err)
		return
	}
	fmt.Println("connect to etcd successful")
	// put操作
	defer cli.Close()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	value := `[{"path":"./nginx.log","topic":"web_log"},{"path":"./mysql.log","topic":"mysql_log"},{"path":"./redis.log","topic":"redis_log"}]`
	//value := `[{"path":"./mysql.log","topic":"mysql_log"},{"path":"./redis.log","topic":"redis_log"}]`
	_, err = cli.Put(ctx, "/logagent/198.18.0.1/collect_config", value)
	cancel()
	if err != nil {
		fmt.Printf("put to etcd failed,err:%v\n", err)
		return
	}
	/*
		//get操作
		ctx, cancel = context.WithTimeout(context.Background(), time.Second)
		resp, err := cli.Get(ctx, "qsb")
		cancel()
		if err != nil {
			fmt.Printf("get from etcd failed,err:%v\n", err)
			return
		}
		for _, ev := range resp.Kvs {
			fmt.Printf("%s:%s\n", ev.Key, ev.Value)
		}

		//watch操作
		//watch用来获取未来更改的通知
		watch := cli.Watch(context.Background(), "qsb")
		for response := range watch {
			for _, event := range response.Events {
				fmt.Printf("event:%v,key:%v,value:%v", event.Type, string(event.Kv.Key),
					string(event.Kv.Value))
			}
		}*/

}
