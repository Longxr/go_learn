package taillog

import (
	"fmt"
	"go_learn/logagent/etcd"
)

var taskMgr *tailLogMgr

// tailLogMgr 日志task管理者
type tailLogMgr struct {
	logEntryList []*etcd.LogEntry
	taskMap map[string]*TailTask
	newConfChan chan []*etcd.LogEntry
}

func Init(logEntryConf []*etcd.LogEntry) {
	taskMgr = &tailLogMgr{
		logEntryList: logEntryConf,	//当前配置信息
		taskMap: make(map[string]*TailTask, 16),
		newConfChan: make(chan []*etcd.LogEntry),	//无缓冲区通道
	}
	for _, logEntry := range logEntryConf {
		fmt.Printf("conf:%v\n", logEntry)
		tailObj := NewTailTask(logEntry.Path, logEntry.Topic)
		key := fmt.Sprintf("%s_%s", logEntry.Path, logEntry.Topic)
		taskMgr.taskMap[key] = tailObj
	}
	go taskMgr.run()
}

func (t * tailLogMgr)run() {
	for {
		select {
		case newConf := <-t.newConfChan:
			for _, conf := range newConf {
				key := fmt.Sprintf("%s_%s", conf.Path, conf.Topic)
				_, ok := t.taskMap[key]
				if ok {
					//已经存在
					continue
				} else {
					//配置新增
					tailObj := NewTailTask(conf.Path, conf.Topic)
					t.taskMap[key] = tailObj
				}
			}
			//配置新增
			//配置删除
			//配置变更
			fmt.Printf("new conf", newConf)
		}
	}
}

func ConfChan() chan<-[]*etcd.LogEntry {
	return taskMgr.newConfChan
}
