package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"io"
	"log"
	"os"
	"sync"
	"time"
)

var (
	Info *log.Logger
	//Warning *log.Logger
	Error *log.Logger
)

func main() {
	logFile, err := os.OpenFile("logs.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("打开日志文件失败：", err)
		panic("打开日志文件失败")
	}

	Info = log.New(io.MultiWriter(os.Stderr, logFile), "Info:", log.Llongfile|log.Lmicroseconds|log.Ldate)
	//Warning = log.New(os.Stdout, "Warning:", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(io.MultiWriter(os.Stderr, logFile), "Error:", log.Llongfile|log.Lmicroseconds|log.Ldate)

	//log.SetOutput(logFile)
	//log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)

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

	//go func() {
	//	for {
	//		get(2)
	//	}
	//}()

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
		Error.Println("producer closed, err:", err)
		return
	}
	defer client.Close()

	index := 1
	for {
		// 构造一个消息
		msg := &sarama.ProducerMessage{}
		msg.Topic = "web_log2"
		txt := fmt.Sprintf("this is a test log %d", index)
		msg.Value = sarama.StringEncoder(txt)

		// 发送消息
		pid, offset, err := client.SendMessage(msg)
		//_, _, err := client.SendMessage(msg)
		if err != nil {
			fmt.Println("send msg failed, err:", err)
			return
		}
		Error.Printf("pid:%v offset:%v msg:%s \n", pid, offset, txt)

		index++
	}
}

var nextOffset = int64(0)

var myLock sync.Mutex

func getOffset() int64 {
	myLock.Lock()
	defer func() {
		nextOffset++
		myLock.Unlock()
	}()
	println(nextOffset)
	return nextOffset
}

func get(ConsumerId int) {
	consumer, err := sarama.NewConsumer([]string{"47.98.135.76:8677"}, nil)
	if err != nil {
		Error.Printf("fail to start consumer, err:%v\n", err)
		return
	}
	partitionList, err := consumer.Partitions("web_log2") // 根据topic取到所有的分区
	if err != nil {
		Error.Printf("fail to get list of partition:err%v\n", err)
		return
	}
	fmt.Println(string(partitionList))

	//partition := 0

	for partition := range partitionList { // 遍历所有的分区
		// 针对每个分区创建一个对应的分区消费者
		pc, err := consumer.ConsumePartition("web_log2", int32(partition), sarama.OffsetNewest)
		if err != nil {
			Error.Printf("failed to start consumer for partition %d,err:%v\n", partition, err)
			return
		}
		defer pc.AsyncClose()
		// 异步从每个分区消费信息
		go func(sarama.PartitionConsumer) {
			for msg := range pc.Messages() {
				Info.Printf("ConsumerId:%d Partition:%d Offset:%d Key:%v Value:%v \n", ConsumerId, msg.Partition, msg.Offset, msg.Key, string(msg.Value))
			}
		}(pc)
	}
}

/**
开多个消费者会有重复消费的问题，只能自己存offest

如果每次都用最新的offest,会有消息丢失的问题。

*/
