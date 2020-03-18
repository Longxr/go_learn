package main

import (
	"logger/mylogger"
)

//全局接口变量
var log mylogger.Logger

func main() {
	id := 10010
	name := "啦啦啦"
	// log = mylogger.NewConsoleLogger("Info")
	log = mylogger.NewFileLogger("Info", "./", "test.log", 10*1024*1024)
	for {
		log.Debug("这是一条Debug日志")
		log.Info("这是一条Info日志")
		log.Warning("这是一条Warning日志")
		log.Error("这是一条Error日志, id:%d, name:%s", id, name)
		log.Fatal("这是一条Fatal日志")
		// time.Sleep(time.Second)
	}
}
