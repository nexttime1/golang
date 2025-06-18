package main

import "github.com/gin-gonic/gin"

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Static("st", "static") //st  代表  就像 GET 的login一样
	r.StaticFile("abc", "static/xtm_skw.txt")
	r.Run(":8080")
}
