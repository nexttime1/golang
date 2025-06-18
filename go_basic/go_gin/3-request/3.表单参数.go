package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode("release")
	r := gin.Default()
	r.POST("/users", func(c *gin.Context) {
		username := c.PostForm("username") //空值  和 没传 分不清
		age, ok := c.GetPostForm("age")

		fmt.Println(username)
		fmt.Println(ok, age)
	})
	r.Run(":8080")
}
