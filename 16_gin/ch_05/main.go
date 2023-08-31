package main

import (
	"awesomeProject/16_gin/ch_05/proto"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.GET("/moreJSON", moreJSON)
	router.GET("/moreProtocolbuf", moreProtocol)
	router.Run(":8083")

	//http://localhost:8083/welecome?firstName=sharui&lastName=123
	//http://localhost:8083/welecome?lastName=123
}

func moreProtocol(context *gin.Context) {
	course := []string{
		"python", "go",
	}
	user := proto.Teacher{
		Name:   "sharui",
		Course: course,
	}
	// HTML的时候使用PureJSON, 会对特殊字符串进行替换。 例如<>
	context.ProtoBuf(http.StatusOK, &user)
}

func moreJSON(context *gin.Context) {
	var msg struct {
		Name    string `json:"user"`
		Message string
		Number  int
	}
	msg.Name = "sharui"
	msg.Message = "消息"
	msg.Number = 100

	context.JSON(http.StatusOK, msg)
}

//{
//"user": "sharui",
//"Message": "消息",
//"Number": 100
//}
