package main

import (
	"fmt"
	"github.com/Shopify/sarama"
)

func main() {
	//初始化一个消费者
	consumer, err := sarama.NewConsumer([]string{"11.2.2.128:9092"}, nil)
	if err != nil {
		fmt.Println("fail ot start consumer,err:", err)
		return
	}
	//	获取指定topic的所有分区
	partitions, err := consumer.Partitions("web_log")
	if err != nil {
		fmt.Println("fail to get partitions,err:", err)
		return
	}
	fmt.Printf("get partitions:%v\n", partitions)
	// 针对每个分区创建一个对应的分区消费者
	for i, partition := range partitions {
		fmt.Printf("i:%#v i.type:%T,partition:%#v partition.type:%T", i, i, partition, partition)
		consumePartition, err := consumer.ConsumePartition("web_log", int32(partition), sarama.OffsetNewest)
		if err != nil {
			fmt.Printf("fail to start ConsumePartition,err:%v", err)
			return
		}
		defer consumePartition.Close() //延迟关闭分区消费者
		//	异步从每个分区消费信息
		go func(partitionConsumer sarama.PartitionConsumer) {
			for msg := range consumePartition.Messages() {
				fmt.Printf("Partition:%d Offset:%d key=%v value:%v \n",
					msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
			}
		}(consumePartition)
	}
	select {}
}
