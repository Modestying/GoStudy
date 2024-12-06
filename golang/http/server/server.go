package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	e := gin.Default()
	e.GET("/get", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello world",
		})
	})
	e.Run("0.0.0.0:8080")
}
