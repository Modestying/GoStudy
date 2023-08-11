package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func HandleIndex(ctx *gin.Context) {
	ctx.String(http.StatusOK, "12")
}

func timerMiddleWare(route string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		startTime := time.Now()
		ctx.Next()
		time.Sleep(time.Second * 2)
		fmt.Println(route, " spend time:", time.Since(startTime).Seconds())
	}
}

func main() {
	engine := gin.Default()
	engine.GET("/index", timerMiddleWare("index"), HandleIndex)
	engine.GET("/usr/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		fmt.Println(name, " ", action)
		c.String(http.StatusOK, name+" is "+action)

	})
	engine.Run(":80")
}
