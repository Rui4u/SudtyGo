package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	goodsGroup := router.Group("goods")
	goodsGroup.GET("list", goodsList)
	goodsGroup.GET(":id", goodDetail)
	goodsGroup.GET(":id/*test", goodDetail)
	router.Run()
}

func goodDetail(context *gin.Context) {
	id := context.Param("id")
	test := context.Param("test")
	context.JSON(http.StatusOK, gin.H{
		"message": "goodDetail",
		"id":      id,
		"*":       test,
	})
}

func goodsList(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "goodList",
	})
}
