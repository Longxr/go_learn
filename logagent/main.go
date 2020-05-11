package main

import (
	"fmt"
	"go_learn/logagent/conf"
	"go_learn/logagent/etcd"
	"go_learn/logagent/kafka"
	"go_learn/logagent/taillog"
	"gopkg.in/ini.v1"
	"sync"
	"time"
)

var cfg = new(conf.AppConf)



// main logagent入口
func main() {
	//加载配置文件
	err := ini.MapTo(cfg, "./conf/config.ini")
	if err != nil {
		fmt.Printf("load config failed, err %v\n", err)
	}

	//初始化kafka连接
	err = kafka.Init([]string{cfg.KafkaConf.Address}, cfg.KafkaConf.ChanMaxSize)
	if err != nil {
		fmt.Printf("init kafka failed, err %v\n", err)
	}
	fmt.Println("init kafka success")
	//初始化etcd
	err = etcd.Init(cfg.EtcdConf.Address, time.Duration(cfg.EtcdConf.Timeout) * time.Second)
	if err != nil {
		fmt.Printf("init etcd failed, err %v\n", err)
	}
	fmt.Println("init etcd success")
	//从etcd中获取日志收集项目配置信息
	err = etcd.PutConf()
	logEntryConf, err := etcd.GetConf(cfg.EtcdConf.Key)
	if err != nil {
		fmt.Printf("etcd.GetConf failed, err %v\n", err)
	}
	//根据配置创建task
	taillog.Init(logEntryConf)
	confChan := taillog.ConfChan()
	var wg sync.WaitGroup
	wg.Add(1)
	//监听配置文件更新
	go etcd.WatchConf(cfg.EtcdConf.Key, confChan)
	wg.Wait()
}
