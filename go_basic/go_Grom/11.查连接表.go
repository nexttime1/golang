package main

import (
	"fmt"
	"time"
)

type UserModel struct {
	ID       uint
	Name     string
	Collects []EssayModel `gorm:"many2many:user_collect_models;joinForeignKey:UserID;JoinReferences:EssayID"`
}

type EssayModel struct {
	ID    uint
	Title string
	// 这里也可以反向引用，根据文章查哪些用户收藏了
}

// UserCollectModel 用户收藏文章表
type UserCollectModel struct {
	UserID     uint       `gorm:"primaryKey"`
	EssayID    uint       `gorm:"primaryKey"`
	UserModel  UserModel  `gorm:"foreignKey:UserID"` //表示当前UserCollectModel 那个字段去关联UserModel
	EssayModel EssayModel `gorm:"foreignKey:EssayID"`
	CreatedAt  time.Time
}

func main() {
	DB.SetupJoinTable(&UserModel{}, "Collects", &UserCollectModel{})
	//err := DB.AutoMigrate(&UserModel{}, &EssayModel{}, &UserCollectModel{})
	//fmt.Println(err)
	//
	//
	//DB.Create(&UserModel{
	//	Name: "薛提猛",
	//	Collects: []EssayModel{
	//		{Title: "我的本科史"},
	//		{Title: "我的研究生史"},
	//	},
	//})

	//如果想要查一个人 的文章  分页  两种
	//第一种
	var u1 UserModel
	DB.Preload("Collects").Take(&u1, "Name = ?", "薛提猛")
	fmt.Println(u1)
	var e1 EssayModel
	DB.Model(&u1).Limit(1).Association("Collects").Find(&e1)
	fmt.Println(e1)

	//第二种 用中间表方式
	var userCollectModel []UserCollectModel

	DB.Preload("UserModel").Preload("EssayModel").Find(&userCollectModel, 1)
	fmt.Println(userCollectModel)
}
