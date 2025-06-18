package main

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.POST("/", func(c *gin.Context) {
		ByteData, err := io.ReadAll(c.Request.Body)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(ByteData))
		c.Request.Body = io.NopCloser(bytes.NewBuffer(ByteData))
		//阅后即焚
		name := c.PostForm("key") //不重新赋值  读不到
		fmt.Println(name)
	})

}
