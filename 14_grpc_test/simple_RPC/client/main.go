package main

import (
	proto2 "awesomeProject/14_grpc_test/simple_RPC/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"time"
)

func main() {
	intercepter := func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		start := time.Now()
		err := invoker(ctx, method, req, reply, cc, opts...)
		fmt.Printf("耗时:%s\n ", time.Since(start))
		return err
	}

	opt := grpc.WithUnaryInterceptor(intercepter)

	conn, error := grpc.Dial("0.0.0.0:8080", grpc.WithTransportCredentials(insecure.NewCredentials()), opt)
	if error != nil {
		panic(error)
	}
	defer conn.Close()

	c := proto2.NewGreeterClient(conn)
	r, error := c.SayHello(context.Background(), &proto2.HelloRequest{Name: "sr"})
	if error != nil {
		panic(error)
	}
	print(r.Message)
}
