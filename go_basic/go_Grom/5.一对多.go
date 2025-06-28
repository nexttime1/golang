package main

import (
	"fmt"
	"gorm.io/gorm"
)

type User struct {
	ID       uint      `gorm:"size:4"`
	Name     string    `gorm:"size:8"`
	Articles []Article // 用户拥有的文章列表
}

type Article struct {
	ID     uint   `gorm:"size:4"`
	Title  string `gorm:"size:16"`
	UserID uint   `gorm:"size:4"` // 属于   这里的类型要和引用的外键类型一致，包括大小
	User   User   // 属于
}

//关于外键命名，外键名称就是关联表名+ID，类型是uint

// 第二种  可以乱起外键
//type User struct {
//	ID       uint      `gorm:"size:4"`
//	Name     string    `gorm:"size:8"`
//	Articles []Article `gorm:"foreignKey:UserID11"` // 用户拥有的文章列表
//}
//
//type Article struct {
//	ID       uint   `gorm:"size:4"`
//	Title    string `gorm:"size:16"`
//	UserID11 uint   `gorm:"foreignKey:UserID11"` // 属于   这里的类型要和引用的外键类型一致，包括大小
//	User     User   // 属于
//}

// 第三种   可以用姓名去代替ID  重写外键关联
//type Customer struct {
//	ID       uint `gorm:"primaryKey"`
//	Name     string
//	Age      int
//	Gender   bool
//	UserInfo UserInfo `gorm:"foreignKey:CustomerID"` // 通过UserInfo可以拿到用户详情信息
//}
//
//type UserInfo struct {
//	CustomerID uint // 外键
//	ID         uint `gorm:"primaryKey"`
//	Addr       string
//	Like       string
//	User       *Customer `gorm:"foreignKey:CustomerID;references:ID"` //解决循环依赖  无法编译   这样指针最大  8字节
//}

func main() {
	DB.AutoMigrate(&User{}, &Article{})

	//添加数据   一下俩表都添加
	DB.Create(&User{
		Name: "宋可为",
		Articles: []Article{{
			Title: "Golang",
		},
			{
				Title: "Python",
			}},
	})
	//创建已经关联的
	DB.Create(&Article{
		Title:  "赵云集",
		UserID: 1,
	})
	//若这个是其他人的
	DB.Create(&Article{
		Title: "张三的书",
		User: User{
			Name: "张三",
		},
	})
	//也可以  先搜在做
	var user User
	DB.Take(&user, 2) // 搜索ID =2的
	DB.Create(&Article{
		Title: "张三的第二本书",
		User:  user,
	})
	//给已有文章绑定用户
	var articles Article
	var u1 User
	DB.Take(&u1, 4)
	DB.Take(&articles, 8) //找到 已有的文章
	DB.Model(&u1).Association("Articles").Append(&articles)

	//给文章绑定用户
	var article01 Article
	var u2 User
	DB.Take(&u2, 5)
	DB.Take(&article01, 9)
	DB.Model(&article01).Association("User").Append(&u2)

	//查询预加载的使用
	var u3 User
	DB.Find(&u3)
	fmt.Println(u3) //{1 宋可为 []}  我想要  articles 也要有

	DB.Preload("Articles").Find(&u3)
	fmt.Println(u3) //{1 宋可为 [{1 Golang 1 {0  []}} {2 Python 1 {0  []}} {5 赵云集 1 {0  []}}]}

	//还可以带条件的  比如 只显示id 大于2 的书籍
	DB.Preload("Articles", "id > ?", 2).Find(&u3)
	fmt.Println(u3)

	//也可以弄个函数
	DB.Preload("Articles", func(db *gorm.DB) *gorm.DB {
		return db.Where("id > ?", 2)
	}).Find(&u3)
	fmt.Println(u3)

	//删除  两种方法
	var u4 User
	DB.Preload("Articles").Find(&u4, 5) //这样就可以 u4.Articles  有数据
	DB.Model(&u4).Association("Articles").Delete(&u4.Articles)

	//第二种删除   删除的更彻底
	var u5 User
	DB.Preload("Articles").Find(&u5, 3) //这样就可以 u4.Articles  有数据
	DB.Select("Articles").Delete(&u5)
}
