package taillog

import (
	"fmt"
	"github.com/hpcloud/tail"
	"go_learn/logagent/kafka"
	"time"
)

// 从日志文件收集日志的模块

// TailTask 日志收集任务
type TailTask struct {
	path string
	topic string
	instance *tail.Tail
}

func NewTailTask(path, topic string) (tailObj *TailTask) {
	tailObj = &TailTask{
		path:path,
		topic:topic,
	}
	tailObj.init()
	return
}

// 初始化日志收集对象
func (t *TailTask)init(){
	config := tail.Config{
		ReOpen: true,
		Follow: true,
		Location: &tail.SeekInfo{Offset: 0, Whence: 2},//记录文件读的位置
		MustExist: false,
		Poll: true,
	}
	var err error
	t.instance, err = tail.TailFile(t.path, config)
	if err != nil {
		fmt.Println("tail file failed, err:", err)
	}
	//开启循环发送日志
	go t.run()
}

func (t* TailTask) run() {
	for {
		// 读取日志
		select {
		case line := <-t.instance.Lines:
			// 发送日志到kafka
			kafka.SendToChan(t.topic, line.Text)
		default:
			time.Sleep(time.Second)
		}
	}
}