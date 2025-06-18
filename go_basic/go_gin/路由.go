package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func UserView(c *gin.Context) {
	path := c.Request.URL.Path
	fmt.Println(path)
	fmt.Println(c.Request.Method)
}

func main() {
	r := gin.Default()
	g := r.Group("/api") //方便路由管理
	g.Use()              //使用中间件
	UserGroup(g)

	NoMid_g := g.Group("/api") //可以一样的，一个是经过中间件  一个不经过
	LoginGroup(NoMid_g)
	r.Run(":8080")
}

func UserGroup(g *gin.RouterGroup) {
	g.GET("/users", UserView)
	g.POST("/users", UserView)
	g.DELETE("/users/", UserView)
}
func LoginGroup(g *gin.RouterGroup) {
	g.GET("/Login", UserView)
	g.POST("/Login", UserView)
	g.DELETE("/Login/", UserView)
}

/*
// 在没有resetful规范正确，表示创建用户，删除用户
/api/user_create
/api/users/create
/api/users/add
/api/add_user
/api/user/delete
/api/user_remove

// 使用resetful规范
GET /api/users  用户列表
POST /api/users  创建用户
PUT /api/users/:id 更新用户信息
DELETE /api/users 批量删除用户
DELETE /api/users/:id 删除单个用户
*/
