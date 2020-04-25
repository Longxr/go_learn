package main

import (
	"fmt"
	"go_learn/logagent/conf"
	"go_learn/logagent/kafka"
	"go_learn/logagent/taillog"
	"gopkg.in/ini.v1"
	"time"
)

var cfg = new(conf.AppConf)

func run() {
	for {
		// 读取日志
		select {
			case line := <-taillog.ReadChan():
				// 发送日志到kafka
				kafka.SendToKafka(cfg.KafkaConf.Topic, line.Text)
			default:
				time.Sleep(time.Second)
		}
	}

}

// main logagent入口
func main() {
	//加载配置文件
	err := ini.MapTo(cfg, "./conf/config.ini")
	if err != nil {
		fmt.Printf("load config failed, err %v\n", err)
	}

	//初始化kafka连接
	err = kafka.Init([]string{cfg.KafkaConf.Address})
	if err != nil {
		fmt.Printf("init kafka failed, err %v\n", err)
	}
	fmt.Println("init kafka success")
	//打开日志文件准备收集
	err = taillog.Init(cfg.TaillogConf.FileName)
	if err != nil {
		fmt.Printf("init taillog failed, err %v\n", err)
	}
	fmt.Println("init taillog success")
	run()
}
