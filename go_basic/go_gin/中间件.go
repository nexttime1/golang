package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func M1(c *gin.Context) {
	fmt.Println("M1  请求部分")
	c.Next() //c.Abort()   那直接返回   也就是  M1  请求部分  M1  响应部分
	fmt.Println("M1  响应部分")
}
func M2(c *gin.Context) {
	fmt.Println("M2  请求部分")
	c.Next() //c.Abort()   那直接返回   也就是  M1  请求部分 M2  请求部分 M2  响应部分  M1  响应部分        不走Home
	fmt.Println("M2  响应部分")
}
func Home(c *gin.Context) {
	fmt.Println("Home")
	c.String(200, "Home")
}

func main() {
	gin.SetMode(gin.DebugMode)
	r := gin.Default()
	r.GET("/", M1, M2, Home) //流程
	/*
		M1  请求部分
		M2  请求部分
		Home
		M2  响应部分
		M1  响应部分
	*/
	r.Run(":8080")
}
