package main

import (
	proto "awesomeProject/14_grpc_test/stream_grpc/proto"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"net"
	"sync"
	"time"
)

const PORT = ":50052"

type server struct {
	*proto.UnimplementedGreeterServer
}

func (*server) GetStream(req *proto.StreamReqData, res proto.Greeter_GetStreamServer) error {
	i := 0
	for {
		i++
		_ = res.Send(&proto.StreamResData{
			Data: fmt.Sprintf("%v", time.Now().Unix()),
		})
		time.Sleep(time.Second)
		if i > 10 {
			break
		}
	}
	return nil
}
func (*server) PostStream(req proto.Greeter_PostStreamServer) error {
	for {
		if a, err := req.Recv(); err != nil {
			fmt.Println(err)
			break
		} else {
			fmt.Println(a.Data)
		}
	}
	return nil
}
func (*server) AllStream(req proto.Greeter_AllStreamServer) error {
	md, ok := metadata.FromIncomingContext(req.Context())
	if ok {
		fmt.Println(md)
	}
	if value, ok := md["name"]; ok {
		fmt.Println(value)
	}

	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for {
			data, _ := req.Recv()
			fmt.Println("收到客户端消息" + data.Data)
		}
	}()

	go func() {
		defer wg.Done()
		for {
			_ = req.Send(&proto.StreamResData{Data: "我是服务器"})
			time.Sleep(time.Second)
		}
	}()
	wg.Wait()
	return nil
}

func main() {
	lis, error := net.Listen("tcp", PORT)
	if error != nil {
		panic(error)
	}
	s := grpc.NewServer()
	proto.RegisterGreeterServer(s, &server{})

	err := s.Serve(lis)
	if err != nil {
		panic(err)
	}
}
