package main

import (
	"awesomeProject/12_rpc/rpc/client_proxy"
	"fmt"
	"net"
	"net/rpc/jsonrpc"
)

/*
json的方式去序列化
*/
func json() {
	// 链接
	conn, err := net.Dial("tcp", "localhost:1234")
	if err != nil {
		fmt.Println(err)
		panic("链接失败")
	}
	//转化为json序列化
	client := jsonrpc.NewClient(conn)

	var replay *string = new(string)
	err = client.Call("HelloService.Hello", "sharui", replay)
	if err != nil {
		fmt.Println(err)
		panic("调用失败")
	}
	fmt.Print(*replay)
}

func main() {
	// 链接 用rpc是因为内部有gob 序列化协议
	client := client_proxy.NewHelloServiceClient("tcp", "localhost:1234")
	var replay *string = new(string)

	err := client.Hello("sharui", replay)
	if err != nil {
		fmt.Println(err)
		panic("调用失败")
	}
	fmt.Print(*replay)
}
