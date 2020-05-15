package main

import (
	"log"
	"net/http"
	"net/rpc"
	. "go_learn/rpc/rpc_http/protocol"
)

func main() {
	rpc.Register(new(Arith)) // 注册rpc服务
	rpc.HandleHTTP()         // 采用http协议作为rpc载体

	if err := http.ListenAndServe(":2333", nil); err != nil {
		log.Fatalln("fatal error: ", err)
	}
}
