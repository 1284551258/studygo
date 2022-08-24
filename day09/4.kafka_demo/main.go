package main

import (
	"fmt"
	"github.com/Shopify/sarama"
)

func main() {

	config := sarama.NewConfig()

	config.Producer.RequiredAcks = sarama.WaitForAll          //发送玩数据需要leader和follow都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 新选出一个partition
	config.Producer.Return.Successes = true                   // 成功交付的消息将在success channel返回

	//构造一个消息
	msg := &sarama.ProducerMessage{}
	msg.Topic = "web_log"
	msg.Value = sarama.StringEncoder("this is a test log")

	//连接kafka
	client, err := sarama.NewSyncProducer([]string{"11.2.2.128:9092"}, config)
	if err != nil {
		fmt.Println("Producer closed,err:", err)
		return
	}
	defer client.Close()
	//发送消息
	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		fmt.Println("Send message failed,err:", err)
		return
	}
	fmt.Printf("Send message successful,pid:%v,offset:%v ", pid, offset)

}
