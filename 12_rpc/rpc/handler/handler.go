package handler

const HelloServiceName = "handle/HelloService"

type NewHelloService struct {
}

func (s *NewHelloService) Hello(request string, replay *string) error {
	*replay = "hello" + request
	return nil
}
