package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	e := gin.Default()
	e.POST("/get", func(c *gin.Context) {
		fmt.Println(c.Request.RemoteAddr)
		c.JSON(200, gin.H{
			"message": "hello world",
		})
	})
	e.Run("0.0.0.0:8080")
}
