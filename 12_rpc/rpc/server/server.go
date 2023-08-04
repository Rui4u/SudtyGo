package main

import (
	"awesomeProject/12_rpc/rpc/handler"
	server_proxy "awesomeProject/12_rpc/rpc/server_proxy"
	"fmt"
	"net"
	"net/rpc"
)

type HelloService struct{}

func (s *HelloService) Hello(request string, reply *string) error {
	*reply = "hello," + request
	return nil
}

func main() {
	// 1. 实例化一个server
	listener, err := net.Listen("tcp", "127.0.0.1:1234")
	if err != nil {
		fmt.Println(err)
		return
	}
	// 2. 注册处理逻辑heander
	_ = server_proxy.RegisterHelloService(&handler.NewHelloService{})

	for { //循环  防止链接后关闭
		// 3.启动服务
		conn, _ := listener.Accept() //当一个新的连接进来的时候
		//rpc.ServeCodec(jsonrpc.NewServerCodec(conn)) // gob序列化转化为json序列化
		go rpc.ServeConn(conn)
	}
}
