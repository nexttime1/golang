package res

import (
	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

var MapData = map[int]string{
	1001: "权限错误",
	1002: "角色错误",
}

func response(c *gin.Context, code int, message string, data any) {
	c.JSON(200, Response{code, message, data})
}

func Ok(c *gin.Context, data any, message string) {
	response(c, 0, message, data)
}

func OkWithData(c *gin.Context, data any) {
	Ok(c, data, "成功")
}

func OkWithMessage(c *gin.Context, message string) {
	response(c, 0, message, gin.H{})
}

func Fail(c *gin.Context, code int, message string, data any) {
	response(c, code, message, data)
}

func FailWithMessage(c *gin.Context, message string) {
	response(c, 1001, message, nil)
}

func FailWithCode(c *gin.Context, code int) {
	msg, ok := MapData[code]
	if !ok {
		msg = "服务错误"
	}
	response(c, code, msg, nil)
}
