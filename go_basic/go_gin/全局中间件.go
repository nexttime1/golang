package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func UsersGroup(g *gin.RouterGroup) {
	g.GET("/users", Home1)
	g.POST("/users", Home1)
	g.PUT("/users", Home1)
	g.DELETE("/users", Home1)

}

type UserInfo struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	gin.SetMode(gin.DebugMode)
	r := gin.Default()
	g := r.Group("/api")
	g.Use(Mid1, Mid2)
	UsersGroup(g)

	r.Run(":8080")
}

func Mid1(c *gin.Context) {
	fmt.Println("Mid1 请求部分")
	c.Set("Mid1", "Mid1") //这周期 都可以拿到
	c.Next()
	fmt.Println("Mid1 响应部分")
}
func Mid2(c *gin.Context) {
	fmt.Println("Mid2 请求部分")
	value, exists := c.Get("Mid1")
	if exists {
		fmt.Println(value)
	}
	var user = UserInfo{Name: "赵子龙", Age: 18}
	c.Set("user", user)
	c.Next()
	fmt.Println("Mid2 响应部分")
}
func Home1(c *gin.Context) {
	fmt.Println("Home")
	value, exists := c.Get("user")
	if exists {
		user, ok := value.(UserInfo)
		if ok {
			fmt.Println(user.Name)
		}
	}
}
