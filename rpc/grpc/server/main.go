package main

import (
	"errors"
	"fmt"
	"net"

	pb "go_learn/rpc/grpc/pb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{}

func (s *server) Multiply(ctx context.Context, req *pb.ArithRequest) (*pb.ArithResponse, error) {
	res := &pb.ArithResponse{Pro:req.GetA() * req.GetB()}
	return res, nil
}

func (s *server) Divide(ctx context.Context, req *pb.ArithRequest) (*pb.ArithResponse, error) {
	if req.GetB() == 0 {
		return nil, errors.New("divide by zero")
	}
	res := &pb.ArithResponse{
		Quo: req.GetA() / req.GetB(),
		Rem: req.GetA() % req.GetB(),
	}
	return res, nil
}

func main() {
	// 监听本地的8972端口
	lis, err := net.Listen("tcp", ":2333")
	if err != nil {
		fmt.Printf("failed to listen: %v", err)
		return
	}
	s := grpc.NewServer() // 创建gRPC服务器
	pb.RegisterArithServiceServer(s, &server{}) // 在gRPC服务端注册服务

	reflection.Register(s) //在给定的gRPC服务器上注册服务器反射服务
	// Serve方法在lis上接受传入连接，为每个连接创建一个ServerTransport和server的goroutine。
	// 该goroutine读取gRPC请求，然后调用已注册的处理程序来响应它们。
	err = s.Serve(lis)
	if err != nil {
		fmt.Printf("failed to serve: %v", err)
		return
	}
}