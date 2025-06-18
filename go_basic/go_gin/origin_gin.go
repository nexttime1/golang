package main

import "github.com/gin-gonic/gin"

type DataList struct {
	Code    int    `json:"code"`
	Massage string `json:"massage"`
	Data    any    `json:"data"`
}

func Index(c *gin.Context) {
	c.JSON(200, DataList{
		Code:    0,
		Massage: "成功",
		Data:    map[string]any{},
	})
}

func main() {
	// 日志  取消
	gin.SetMode("release")

	//1.初始化
	r := gin.Default() //引擎
	// 2. 挂载路由
	r.GET("/index", Index)
	//3. 绑定端口  运行
	r.Run("127.0.0.1:8080")
}
