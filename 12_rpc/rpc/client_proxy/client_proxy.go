package client_proxy

import "net/rpc"

type HelloServiceStub struct {
	*rpc.Client
}

func NewHelloServiceClient(protocol, address string) HelloServiceStub {
	conn, error := rpc.Dial(protocol, address)
	if error != nil {
		panic("connect error !")
	}
	return HelloServiceStub{conn}
}

func (s *HelloServiceStub) Hello(request string, replay *string) error {
	*replay = "hello," + request
	return nil
}
