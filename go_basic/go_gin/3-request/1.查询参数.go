package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

//http://127.0.0.1:8080/?name=xtm&age=22&key=123&key=222

func main() {
	gin.SetMode("release")
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		name := c.Query("name")            //查询路径
		age := c.DefaultQuery("age", "18") // 有默认值
		KeyList := c.QueryArray("key")     //多个key  key=22&key=23  接受到一个list
		fmt.Println(name, age, KeyList)
	})
	r.Run(":8080")

}
