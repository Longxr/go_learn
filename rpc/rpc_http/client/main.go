package main

import (
	"fmt"
	"log"
	"net/rpc"
	. "go_learn/rpc/rpc_http/protocol"
)

func main() {
	conn, err := rpc.DialHTTP("tcp", "127.0.0.1:2333")
	if err != nil {
		log.Fatalln("dailing error: ", err)
	}

	req := ArithRequest{7, 3}
	var res ArithResponse

	err = conn.Call("Arith.Multiply", req, &res) // 乘法运算
	if err != nil {
		log.Fatalln("arith error: ", err)
	}
	fmt.Printf("%d * %d = %d\n", req.A, req.B, res.Pro)

	err = conn.Call("Arith.Divide", req, &res)
	if err != nil {
		log.Fatalln("arith error: ", err)
	}
	fmt.Printf("%d / %d, quo is %d, rem is %d\n", req.A, req.B, res.Quo, res.Rem)
}
