package main

import (
	"awesomeProject/14_grpc_test/proto"
	"context"
	"google.golang.org/grpc"
	"net"
)

type ServerHH struct {
	*proto.UnimplementedGreeterServer
}

func (c *ServerHH) SayHello(ctx context.Context, request *proto.HelloRequest) (*proto.HelloReply, error) {
	return &proto.HelloReply{
		Message: "hello" + request.Name,
	}, nil
}

func main() {
	g := grpc.NewServer()
	proto.RegisterGreeterServer(
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
