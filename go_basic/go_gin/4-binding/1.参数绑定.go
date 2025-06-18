package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.POST("/", func(c *gin.Context) { //正常  param 参数
		type User struct {
			Name string `form:"name"`
			Age  int    `form:"age"`
		}
		var user User
		err := c.ShouldBindQuery(&user)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(user.Name)
		fmt.Println("年龄是：", user.Age)
	})
	r.GET("users/:name/:address", func(c *gin.Context) { //路径参数
		type Person struct {
			Name    string `form:"name"`
			Address string `form:"address"`
		}
		var person Person
		err := c.ShouldBindUri(&person)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("地址为", person.Address)
	})
	r.POST("form", func(c *gin.Context) { //接收表单
		type Person struct {
			Name  string `form:"name"`
			Grade int    `form:"grade"`
		}
		var person Person
		err := c.ShouldBind(&person)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("成绩为：", person.Grade)

	})
	r.POST("/json", func(c *gin.Context) { //json  格式
		type Person struct {
			Name  string `json:"name" binding:"oneof=xtm zhangsan"` //必须是这两个中的一个
			Psd   string `json:"psd" binding:"required,e"`
			RePsd string `json:"rePsd" binding:"required,eqfield=Psd" ` //必须一样
		}
		var person Person
		err := c.ShouldBindJSON(&person)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(" json 你的名字是", person.Name)
	})
	r.POST("/header", func(c *gin.Context) {
		type Person struct {
			Name      string `Header:"Name"`
			Age       int    `Header:"Age"`
			UserAgent string `Header:"User-Agent"`
		}
		var person Person
		err := c.ShouldBindHeader(&person)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("header   Age： ", person.Age)
		fmt.Println("header   UserAgent： ", person.UserAgent)
	})

	r.Run(":8080")
}

/*
// 不能为空，并且不能没有这个字段
required： 必填字段，如：binding:"required"

// 针对字符串的长度
min 最小长度，如：binding:"min=5"
max 最大长度，如：binding:"max=10"
len 长度，如：binding:"len=6"

// 针对数字的大小
eq 等于，如：binding:"eq=3"
ne 不等于，如：binding:"ne=12"
gt 大于，如：binding:"gt=10"
gte 大于等于，如：binding:"gte=10"
lt 小于，如：binding:"lt=10"
lte 小于等于，如：binding:"lte=10"

// 针对同级字段的
eqfield 等于其他字段的值，如：PassWord string `binding:"eqfield=Password"`
nefield 不等于其他字段的值
startwith 字符串前缀
endswith  字符串后缀

- 忽略字段，如：binding:"-"
dive   校验针对每一个元素   比喻 list  存放Ip
ip
ipv4
ipv6


*/
