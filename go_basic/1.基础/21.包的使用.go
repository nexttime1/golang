package main

import (
	"github.com/gin-gonic/gin"
	"golang/tools/users"
)

func main() {
	roman := scorce.Roman{Name: "<UNK>", Age: 10}
	roman.GetName()
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, Gin!")
	})
	r.Run(":8080")
}
