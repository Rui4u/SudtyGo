package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
go path 模式

go env
1. 已经将新建的代码到gopath目录下的scr下
2. go env -w go111module=off

import 查找循序
1. gopath/src下是否有
2. goroot/src目录下找

不做包管理
*/

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
