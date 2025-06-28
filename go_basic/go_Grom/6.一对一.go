package main

import (
	"fmt"
	"gorm.io/gorm"
)

type Customer struct {
	ID       uint `gorm:"primaryKey"`
	Name     string
	Age      int
	Gender   bool
	UserInfo UserInfo `gorm:"foreignKey:CustomerID"` // 通过UserInfo可以拿到用户详情信息
}

type UserInfo struct {
	CustomerID uint // 外键
	ID         uint `gorm:"primaryKey"`
	Addr       string
	Like       string
	User       *Customer `gorm:"foreignKey:CustomerID;references:ID"` //解决循环依赖  无法编译   这样指针最大  8字节
}

func main() {
	DB := DB.Session(&gorm.Session{
		Logger: mysqlLogger,
	})
	//DB.AutoMigrate(&Customer{}, UserInfo{}) //创建两个表
	//添加
	//DB.Create(&Customer{
	//	Name:   "薛提猛",
	//	Age:    18,
	//	Gender: true,
	//	UserInfo: UserInfo{
	//		Addr: "山东省",
	//		Like: "打篮球",
	//	},
	//})
	////查看
	//var c1 Customer
	//DB.Preload("UserInfo").Find(&c1, 1)
	//fmt.Println(c1)
	//
	//var userinfo UserInfo
	//DB.Preload("User").Find(&userinfo)
	//fmt.Println(userinfo)
	////删除
	var customer Customer
	DB.Preload("UserInfo").Take(&customer, 1)
	fmt.Println(customer)
	DB.Select("UserInfo").Delete(&customer)

}
