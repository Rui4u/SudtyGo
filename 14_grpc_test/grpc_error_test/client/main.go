package main

import (
	proto "awesomeProject/14_grpc_test/grpc_token_auth_test/proto"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type customRPCredenials struct {
}

func (c customRPCredenials) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"appid":  "0101010101",
		"appkey": "sharui",
	}, nil
}

// RequireTransportSecurity indicates whether the credentials requires
// transport security.
func (c customRPCredenials) RequireTransportSecurity() bool {
	return false
}
func main() {
	// 方式1
	//intercepter := func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	//	start := time.Now()
	//	md := metadata.New(map[string]string{
	//		"appid":  "0101010101",
	//		"appkey": "sharui",
	//	})
	//	ctx = metadata.NewOutgoingContext(ctx, md)
	//	err := invoker(ctx, method, req, reply, cc, opts...)
	//	fmt.Printf("耗时:%s\n ", time.Since(start))
	//	return err
	//}
	//
	//opt := grpc.WithUnaryInterceptor(intercepter)

	//方式2
	opt := grpc.WithPerRPCCredentials(customRPCredenials{})

	conn, error := grpc.Dial("0.0.0.0:8080", grpc.WithTransportCredentials(insecure.NewCredentials()), opt)
	if error != nil {
		panic(error)
	}
	defer conn.Close()

	c := proto.NewGreeterClient(conn)
	r, error := c.SayHello(context.Background(), &proto.HelloRequest{Name: "sr", Id: 10000, Email: "123@lagou.com"})
	if error != nil {
		panic(error)
	}
	print(r.Message)
}
