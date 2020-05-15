package main

import (
	"fmt"
	. "go_learn/rpc/jsonrpc/protocol"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	rpc.Register(new(Arith)) // 注册rpc服务

	lis, err := net.Listen("tcp", ":2333")
	if err != nil {
		log.Fatalln("fatal error: ", err)
	}

	fmt.Println("start connection")

	for {
		conn, err := lis.Accept() // 接收客户端连接请求
		if err != nil {
			continue
		}

		go func(conn net.Conn) { // 并发处理客户端请求
			fmt.Println("new client in coming")
			jsonrpc.ServeConn(conn)
		}(conn)
	}
}
