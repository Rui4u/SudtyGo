package main

import (
	proto "awesomeProject/14_grpc_test/grpc_error_test/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"time"
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

	//方式2
	opt := grpc.WithPerRPCCredentials(customRPCredenials{})

	conn, error := grpc.Dial("0.0.0.0:8080", grpc.WithTransportCredentials(insecure.NewCredentials()), opt)

	if error != nil {
		if st, ok := status.FromError(error); ok {
			fmt.Println(st.Message())
			fmt.Println(st.Code())
		}

	}
	defer conn.Close()

	c := proto.NewGreeterClient(conn)

	ctx, _ := context.WithTimeout(context.Background(), time.Second*4)

	r, error := c.SayHello(ctx, &proto.HelloRequest{Name: "sr"})
	if error != nil {
		panic(error)
	}
	print(r.Message)
}
