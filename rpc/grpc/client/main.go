package main

import (
"context"
"fmt"
"google.golang.org/grpc"
pb "go_learn/rpc/grpc/pb"
)

func main() {
	// 连接服务器
	conn, err := grpc.Dial(":2333", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("faild to connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewArithServiceClient(conn)
	// 调用服务端的SayHello
	req := &pb.ArithRequest{A:9, B:2}
	res, err := c.Multiply(context.Background(), req)
	if err != nil {
		fmt.Printf("arith error:: %v", err)
	}
	fmt.Printf("%d * %d = %d\n", req.GetA(), req.GetB(), res.GetPro())

	res, err = c.Divide(context.Background(), req)
	if err != nil {
		fmt.Printf("arith error:: %v", err)
	}
	fmt.Printf("%d / %d, quo is %d, rem is %d\n", req.A, req.B, res.Quo, res.Rem)
}