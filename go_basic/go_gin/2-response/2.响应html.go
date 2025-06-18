package main

import "github.com/gin-gonic/gin"

func main() {
	gin.SetMode("release")
	r := gin.Default()
	r.LoadHTMLGlob("2-response/templates/*") // 加载一堆文件  相对路径
	//r.LoadHTMLFiles("templates/index.html")  //  加载一个
	r.GET("/", func(c *gin.Context) {
		//HTML  三个参数  第一个状态码   第二个 哪一个文件  必须提前Load进来  第三个传参  .title 可用
		c.HTML(200, "index.html", map[string]any{
			"title": "婚礼",
		})
	})
	r.Run(":8080")
}
