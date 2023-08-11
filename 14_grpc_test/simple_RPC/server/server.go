package main

import (
	proto2 "awesomeProject/14_grpc_test/HttpTest/proto"
	"context"
	"google.golang.org/grpc"
	"net"
)

type ServerHH struct {
	*proto2.UnimplementedGreeterServer
}

func (c *ServerHH) SayHello(ctx context.Context, request *proto2.HelloRequest) (*proto2.HelloReply, error) {
	return &proto2.HelloReply{
		Message: "hello" + request.Name,
	}, nil
}

func main() {
	g := grpc.NewServer()
	proto2.RegisterGreeterServer(
		g,
		&ServerHH{},
	)
	lis, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		panic("failed to listen" + err.Error())
	}
	_ = g.Serve(lis)
	if err != nil {
		panic("fail to start")
	}
}
