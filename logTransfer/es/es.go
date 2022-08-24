package es

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"strings"
	"time"
)

var (
	client *elastic.Client
	ch     chan *LogData
)

type LogData struct {
	Topic string `json:"topic"`
	Data  string `json:"data"`
}

func Init(address string, ch_size, nums int) (err error) {
	//初始化client
	if !strings.HasPrefix(address, "http://") {
		address = "http://" + address
	}
	client, err = elastic.NewClient(elastic.SetURL(address))
	if err != nil {
		fmt.Printf("connect to es faild,err:%s", err)
		return
	}
	fmt.Println("connect to es successful!")
	ch = make(chan *LogData, ch_size)
	for i := 0; i < nums; i++ {
		go SedToEs()
	}
	return
}

func SedToEsChan(msg *LogData) {
	ch <- msg
}
func SedToEs() {
	for true {
		select {
		case msg := <-ch:
			response, err := client.Index().Index(msg.Topic).BodyJson(msg).Do(context.Background())
			if err != nil {
				fmt.Printf("set %v to es failed,err:%v", msg.Data, err)
			}
			fmt.Printf("Indexed user %s to index %s,type %s\n", response.Id, response.Index, response.Type)
		default:
			time.Sleep(time.Second)
		}
	}

}
