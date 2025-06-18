package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

func main() {
	gin.SetMode("release")
	r := gin.Default()
	r.POST("/users", func(c *gin.Context) {
		fileHeader, err := c.FormFile("file")
		if err != nil {
			panic(err)
			return
		}
		fmt.Println(fileHeader.Filename) //名称
		fmt.Print(fileHeader.Size)       //字节数
		file, err := fileHeader.Open()
		if err != nil {
			panic(err)
			return
		}
		var buf []byte
		buf, err = io.ReadAll(file)

		os.WriteFile("xx.jpg", buf, 0666)
		//可以省略成一句话
		//c.SaveUploadedFile(fileHeader,"xx.jpg")
	})
	r.Run(":8080")
}
