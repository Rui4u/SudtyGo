package main

import (
	"awesomeProject/12_rpc/rpc/proto"
	"fmt"
	"github.com/goccy/go-json"
	"google.golang.org/protobuf/proto"
)

type JsonModel struct {
	Name string `json:"name"`
}

func main() {
	req := HelloWorld.HelloRequest{Name: "sharui"}
	rsp, _ := proto.Marshal(&req)

	reqq := HelloWorld.HelloRequest{}
	_ = proto.Unmarshal(rsp, &reqq)
	//
	fmt.Println(reqq.Name)

	fmt.Println(string(rsp))

	jsonReq := JsonModel{"hallo"}
	jsonRsp, _ := json.Marshal(jsonReq)
	fmt.Println(string(jsonRsp))
}
