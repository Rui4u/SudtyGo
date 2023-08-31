package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Person struct {
	ID   string `uri:"id" binding:"required,uuid"`
	Name string `uri:"name" binding:"required""`
}

func main() {
	router := gin.Default()
	router.GET("/:name/:id", func(c *gin.Context) {
		var persion Person
		if err := c.ShouldBindUri(&persion); err != nil {
			c.Status(404)
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"name": persion.Name,
			"id":   persion.ID,
		})
	})
	router.Run()
}
