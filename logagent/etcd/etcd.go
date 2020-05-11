package etcd

import (
	"context"
	"encoding/json"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
)

var (
	cli *clientv3.Client
)

// 需要收集的日志的配置信息
type LogEntry struct {
	Path string `json:"path"`	//日志存放路径
	Topic string `json:"topic"`	//日志发往kafka的topic
}

// Init 初始化etcd函数
func Init(addr string, timeout time.Duration)(err error) {
	cli, err = clientv3.New(clientv3.Config{
		Endpoints: []string{addr},
		DialTimeout: timeout,
	})
	if err != nil {
		//handle error!
		fmt.Printf("connect to etcd failed, err:%v\n", err)
		return
	}
	return
}

// GetConf 从etcd中根据key获取配置
func GetConf(key string) (logConf []*LogEntry, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.Get(ctx, key)
	cancel()
	if err != nil {
		//handle error!
		fmt.Printf("get from etcd failed, err:%v\n", err)
		return
	}
	for _, ev := range resp.Kvs {
		//fmt.Printf("%s:%s\n", ev.Key, ev.Value)
		err = json.Unmarshal(ev.Value, &logConf)
		if err != nil {
			//handle error!
			fmt.Printf("unmarshal etcd value failed, err:%v\n", err)
			return
		}
	}
	return
}

// PutConf 往etcd中写入配置
func PutConf() (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	value := `[{"path":"c:/tmp/nginx.log","topic":"web_log"},{"path":"c:/tmp/redis.log","topic":"redis_log"},{"path":"d:/tmp/mysql.log","topic":"mysql_log"}]`
	_, err = cli.Put(ctx, "/logagent/collect_config", value)
	cancel()
	if err != nil {
		//handle error!
		fmt.Printf("put to etcd failed, err:%v\n", err)
		return
	}
	return
}
