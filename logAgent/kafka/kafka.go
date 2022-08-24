package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
	"time"
)

//定义一个全局的kafkaclient
var client sarama.SyncProducer

type logData struct {
	topic string
	data  string
}

var logDataChan chan *logData

func Init(addr []string, chanMaxSize int) (err error) {
	//初始化配置文件
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll          //发送玩数据需要leader和follow都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 新选出一个partition
	config.Producer.Return.Successes = true
	//连接kafka
	client, err = sarama.NewSyncProducer(addr, config)
	if err != nil {
		fmt.Println("Producer closed,err:", err)
		return
	}
	logDataChan = make(chan *logData, chanMaxSize)
	//开启后台的goroutine从通道获取数据发到kafka
	go sendToKafka(logDataChan)
	return
}

func SendToChan(topic, data string) {
	lg := &logData{
		topic: topic,
		data:  data,
	}
	logDataChan <- lg
}

func sendToKafka(logDataChan chan *logData) {
	//构造一个消息
	for true {
		select {
		case ld := <-logDataChan:
			msg := &sarama.ProducerMessage{}
			msg.Topic = ld.topic
			msg.Value = sarama.StringEncoder(ld.data)
			pid, offset, err := client.SendMessage(msg)
			if err != nil {
				fmt.Println("Send message failed,err:", err)
			}
			fmt.Printf("Send message successful,pid:%v,offset:%v data:%v\n", pid, offset, ld.data)
		default:
			time.Sleep(time.Millisecond * 50)
		}
	}
}
