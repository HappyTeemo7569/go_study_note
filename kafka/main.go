package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"log"
	"time"
)

func main() {
	go func() {
		for {
			put()
		}
	}()

	go func() {
		for {
			get(1)
		}
	}()

	go func() {
		for {
			get(2)
		}
	}()

	for {
		time.Sleep(1 * time.Second)
	}
}

func put() {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll          // 发送完数据需要leader和follow都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 新选出一个partition
	config.Producer.Return.Successes = true                   // 成功交付的消息将在success channel返回

	// 连接kafka
	client, err := sarama.NewSyncProducer([]string{"47.98.135.76:8677"}, config)
	if err != nil {
		fmt.Println("producer closed, err:", err)
		return
	}
	defer client.Close()

	index := 1
	for {
		// 构造一个消息
		msg := &sarama.ProducerMessage{}
		msg.Topic = "web_log"
		txt := fmt.Sprintf("this is a test log %d", index)
		msg.Value = sarama.StringEncoder(txt)

		// 发送消息
		//pid, offset, err := client.SendMessage(msg)
		_, _, err := client.SendMessage(msg)
		if err != nil {
			fmt.Println("send msg failed, err:", err)
			return
		}
		//fmt.Printf("pid:%v offset:%v\n", pid, offset)

		index++
	}
}

func get(ConsumerId int) {
	consumer, err := sarama.NewConsumer([]string{"47.98.135.76:8677"}, nil)
	if err != nil {
		fmt.Printf("fail to start consumer, err:%v\n", err)
		return
	}
	partitionList, err := consumer.Partitions("web_log") // 根据topic取到所有的分区
	if err != nil {
		fmt.Printf("fail to get list of partition:err%v\n", err)
		return
	}
	fmt.Println(string(partitionList))

	for partition := range partitionList { // 遍历所有的分区
		// 针对每个分区创建一个对应的分区消费者
		pc, err := consumer.ConsumePartition("web_log", int32(partition), sarama.OffsetNewest)
		if err != nil {
			fmt.Printf("failed to start consumer for partition %d,err:%v\n", partition, err)
			return
		}
		defer pc.AsyncClose()
		// 异步从每个分区消费信息
		go func(sarama.PartitionConsumer) {
			for msg := range pc.Messages() {
				logs, _ := fmt.Printf("ConsumerId:%d Partition:%d Offset:%d Key:%v Value:%v \n", ConsumerId, msg.Partition, msg.Offset, msg.Key, string(msg.Value))
				log.Println(logs)
			}
		}(pc)
	}
}
