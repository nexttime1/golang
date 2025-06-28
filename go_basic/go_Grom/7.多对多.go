package main

import (
	"fmt"
	"gorm.io/gorm"
)

type Tag struct {
	ID     uint
	Name   string
	essays []Essay `gorm:"many2many:essay_tags;"` // 用于反向引用
}

// Essay Tag 多对多的关系  这种关系需要一个中间表来关联两个模型，而这个中间表的名称就是 essay_tags
type Essay struct {
	ID    uint
	Title string
	Tags  []Tag `gorm:"many2many:essay_tags;"`
}

func main() {
	DB := DB.Session(&gorm.Session{
		Logger: mysqlLogger,
	})
	//创建表  三个表  两个表加一个关系
	DB.AutoMigrate(&Tag{}, &Essay{})

	//添加  什么都没有  两个纯添加
	DB.Create(&Essay{
		Title: "计算机应用",
		Tags: []Tag{
			{Name: "python"},
			{Name: "go"},
			{Name: "java"},
		},
	})
	//如果有 tag  去添加 Essay
	var tags []Tag
	DB.Find(&tags, 1)
	DB.Create(&Essay{
		Title: "计算机软件工程",
		Tags:  tags,
	})
	//查询
	var essays []Essay
	DB.Preload("Tags").Find(&essays, 1)
	fmt.Println(essays)

	//多对多关系 的更新  先删除原有的  在添加
	var e1 Essay
	DB.Preload("Tags").Find(&e1, 1)
	fmt.Println(e1)
	//Association.Delete() 专门用于删除关系
	DB.Model(&e1).Association("Tags").Delete(e1.Tags) //只删除关系表

	//var tags []Tag
	DB.Find(&tags, 1)
	DB.Model(&e1).Association("Tags").Append(&tags)

	//多对多关系的替换
	var e2 Essay
	DB.Preload("Tags").Find(&e2, 1)
	var t1 []Tag
	DB.Find(&t1, 2)
	DB.Model(&e2).Association("Tags").Replace(&t1)
}
