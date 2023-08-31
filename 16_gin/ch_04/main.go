package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.GET("/welecome", welecome)
	router.POST("/formPost", fromPost)
	router.Run(":8083")

	//http://localhost:8083/welecome?firstName=sharui&lastName=123
	//http://localhost:8083/welecome?lastName=123
}

func fromPost(context *gin.Context) {
	message := context.PostForm("message")
	context.JSON(http.StatusOK, gin.H{
		"message": message,
	})
}

func welecome(context *gin.Context) {
	firstName := context.DefaultQuery("firstName", "defaultName")
	lastName := context.Query("lastName")

	context.JSON(http.StatusOK, gin.H{
		"firstName": firstName,
		"lastName":  lastName,
	})
}
