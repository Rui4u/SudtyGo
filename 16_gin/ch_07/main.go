package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func MYLogger() gin.HandlerFunc {
	return func(context *gin.Context) {
		t := time.Now()
		context.Next()
		end := time.Since(t)
		fmt.Println(end)
	}
}

func main() {
	router := gin.New()
	router.Use(MYLogger(), gin.Recovery())

}
