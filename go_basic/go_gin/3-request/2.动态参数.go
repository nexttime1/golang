package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode("release")
	r := gin.Default()
	r.GET("/users/:name", func(c *gin.Context) {
		username := c.Param("name")
		fmt.Println(username)
	})

	r.GET("/users/:name/:age", func(c *gin.Context) {
		username := c.Param("name")
		age := c.Param("age")
		fmt.Println(username, age)
	})
	r.Run(":8080")
}
