package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc/metadata"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	proto "awesomeProject/14_grpc_test/stream_grpc/proto"
)

func main() {

	conn, error := grpc.Dial("0.0.0.0:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if error != nil {
		panic(error)
	}
	defer conn.Close()

	c := proto.NewGreeterClient(conn)

	//res, _ := c.GetStream(context.Background(), &proto.StreamReqData{Data: "sharui"})
	//for {
	//	a, error := res.Recv()
	//	if error != nil {
	//		fmt.Println(error)
	//		break
	//	}
	//	println(a.Data)
	//}
	//
	////客户端流模式
	//puts, _ := c.PostStream(context.Background())
	//i := 0
	//for {
	//	i++
	//	puts.Send(&proto.StreamReqData{Data: fmt.Sprintf("sharui%d", i)})
	//	time.Sleep(time.Second)
	//	if i > 10 {
	//		break
	//	}
	//}

	// 双向流模式

	md := metadata.New(map[string]string{"name": "sharui"})
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	all, _ := c.AllStream(ctx)
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		for {
			all.Send(&proto.StreamReqData{Data: "发送一个数据"})
			time.Sleep(time.Second)
		}
	}()

	go func() {
		defer wg.Done()
		for {
			a, err := all.Recv()
			if err != nil {
				fmt.Println(err)
				break
			}
			fmt.Println("收到服务器消息" + a.Data)
		}
	}()

	wg.Wait()

}
