package main

import (
	"github.com/gin-gonic/gin"
	"go_gin/2-response/res"
)

func main() {
	gin.SetMode("release")
	r := gin.Default()
	r.GET("/login", func(c *gin.Context) {
		//c.JSON(200, gin.H{"code": 0, "message": "ok", "data": gin.H{}}) //type H map[string]any  就是一个map
		res.OkWithMessage(c, "登录成功")
	})
	r.GET("/users", func(c *gin.Context) {
		res.OkWithData(c, map[string]any{
			"username": "薛提猛",
		})
	})
	r.POST("/users", func(c *gin.Context) {
		res.FailWithMessage(c, "参数错误")
	})

	r.Run(":8080")
}
