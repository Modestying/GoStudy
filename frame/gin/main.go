package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
	"github.com/gin-contrib/timeout"
	"github.com/gin-gonic/gin"
)

func timeoutMiddleware() gin.HandlerFunc {
	return timeout.New(
		timeout.WithTimeout(3000*time.Millisecond),
		timeout.WithHandler(func(c *gin.Context) {
			c.Next()
		}),
		timeout.WithResponse(testResponse),
	)
}
func HandleIndex(ctx *gin.Context) {
	funcCtx, cancelFunc := context.WithTimeout(context.Background(), time.Second*10)
	funcCtx = context.WithValue(funcCtx, "res", 1)
	defer cancelFunc()
	var a int
	go DoSth(funcCtx, &a)
	select {
	case <-funcCtx.Done():
		ctx.String(http.StatusGatewayTimeout, "timeout")
	}

	ctx.String(http.StatusOK, fmt.Sprintf("%d", a))

}

func DoSth(ctx context.Context, val *int) int {
	time.Sleep(time.Second * 5)
	*val = 111
	return 5
}
func timerMiddleWare(route string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		startTime := time.Now()
		ctx.Next()
		//time.Sleep(time.Second * 2)
		fmt.Println(route, " spend time:", time.Since(startTime).Seconds())
	}
}

type User struct {
	Name     string `json:"name" form:"name"`
	Password string `json:"passwd" form:"passwd"`
}

func main() {
	engine := gin.Default()
	engine.GET("/index", timerMiddleWare("index"), HandleIndex)

	engine.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "GIN version 1")
	})
	engine.GET("/json", func(ctx *gin.Context) {
		// 返回json 数据
		// 1.使用map自定义数据
		// 2.通过结构体返回
		// 3.使用gin.H{}，实质也是map
		usr := User{
			Name: "this is  json data",
		}
		ctx.JSON(http.StatusOK, usr)
	})
	// query?name=sd&passwd=123
	engine.GET("/query", func(ctx *gin.Context) {
		fmt.Println("/query:", ctx.Query("name")+" "+ctx.Query("passwd"))
	})

	// rest 风格api
	engine.GET("/usr/:name/:action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		fmt.Println(name, " ", action)
		c.String(http.StatusOK, name+" is "+action)

	})
	// curl -d 'name=ax' -d 'passwd=123' -X POST http://localhost:80/login
	engine.POST("/login", func(ctx *gin.Context) {
		fmt.Println(ctx.PostForm("name"))
		fmt.Println(ctx.PostForm("passwd"))

		usr := User{}
		if ctx.ShouldBind(&usr) != nil {
			panic("xx")
		}
		fmt.Println(usr)

	})
	engine.POST("/auth", func(ctx *gin.Context) {

	})
	engine.Run(":8085")
}
