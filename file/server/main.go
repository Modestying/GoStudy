package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	//限制上传最大尺寸
	r.MaxMultipartMemory = 100 << 20
	r.POST("/upload", func(c *gin.Context) {
		file, err := c.MultipartForm()
		if err != nil {
			c.String(500, "上传图片出错")
		}
		for x, _ := range file.File {
			fmt.Printf("文件名:%s\n", x)
		}
		c.SaveUploadedFile(file.File["front"][0], file.File["front"][0].Filename)
		c.String(http.StatusOK, file.File["front"][0].Filename+" 上传成功")
	})
	r.Run("0.0.0.0:8976")
}
