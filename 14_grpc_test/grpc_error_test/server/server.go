package main

import (
	proto2 "awesomeProject/14_grpc_test/grpc_error_test/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"net"
)

type ServerHH struct {
	*proto2.UnimplementedGreeterServer
}

func (c *ServerHH) SayHello(ctx context.Context, request *proto2.HelloRequest) (*proto2.HelloReply, error) {

	return nil, status.Error(codes.NotFound, "没找到")
	return &proto2.HelloReply{
		Message: "hello" + request.Name,
	}, nil
}

func main() {

	interceptor := func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return resp, status.Error(codes.Unauthenticated, "无token认证信息")
			fmt.Println("get metadata error")
		}

		var (
			appkey string
			appid  string
		)

		if va1, ok := md["appid"]; ok {
			appid = va1[0]
		}

		if va1, ok := md["appkey"]; ok {
			appkey = va1[0]
		}

		if appkey != "sharui" || appid != "0101010101" {
			return resp, status.Error(codes.Unauthenticated, "无token认证信息")
		}

		fmt.Println("接收一个新请求")

		res, err := handler(ctx, req)
		fmt.Println("请求已经完成")

		return res, err
	}

	opt := grpc.UnaryInterceptor(interceptor) //生成拦截器
	g := grpc.NewServer(opt)

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
