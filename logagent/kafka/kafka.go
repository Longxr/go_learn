package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
	"time"
)

// 往kafka写日志的模块

type logData struct {
	topic string
	data string
}

var (
	client sarama.SyncProducer
	logDataChan chan *logData
)

// Init 初始化kafka连接
func Init(addrs []string, maxSize int)(err error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll          // 发送完数据需要leader和follow都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 新选出一个partition
	config.Producer.Return.Successes = true                   // 成功交付的消息将在success channel返回

	// 连接kafka
	client, err = sarama.NewSyncProducer(addrs, config)
	if err != nil {
		fmt.Println("producer closed, err:", err)
		return
	}
	//初始化通道
	logDataChan = make(chan *logData, maxSize)
	//后台开始等待数据
	go sendToKafka()
	return
}

// 外部调用函数，日志放到chan中
func SendToChan(topic, data string) {
	msg := &logData{
		topic:topic,
		data:data,
	}
	logDataChan<-msg
}

// 从通道取一条日志，发送到kafka
func sendToKafka() {
	for {
		select {
			case ld := <-logDataChan:
				// 构造一个消息
				msg := &sarama.ProducerMessage{}
				msg.Topic = ld.topic
				msg.Value = sarama.StringEncoder(ld.data)
				// 发送消息
				pid, offset, err := client.SendMessage(msg)
				if err != nil {
					fmt.Println("send msg failed, err:", err)
					return
				}
				fmt.Printf("pid:%v offset:%v\n", pid, offset)
		default:
			time.Sleep(time.Millisecond*50)
		}
	}
}

