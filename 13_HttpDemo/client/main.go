package main

import (
	"encoding/json"
	"fmt"
	"github.com/kirinlabs/HttpRequest"
)

type ResponseData struct {
	Data int `json:"data"`
}

func add(a, b int) int {
	req := HttpRequest.NewRequest()
	res, _ := req.Get(fmt.Sprintf("http://127.0.0.1:8000/%s?a=%d&b=%d", "add", a, b))
	body, _ := res.Body()
	fmt.Println(string(body))
	responseData := ResponseData{}
	json.Unmarshal(body, &responseData)
	return responseData.Data
}

func main() {
	result := add(4, 2)
	fmt.Println(result)

}
