package main

import (
	proto2 "awesomeProject/14_grpc_test/HttpTest/proto"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, error := grpc.Dial("0.0.0.0:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
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
