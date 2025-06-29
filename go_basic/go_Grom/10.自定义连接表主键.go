package main

import "time"

type ArticleModel struct {
	ID    uint
	Title string
	Tags  []TagModel `gorm:"many2many:article_tags;joinForeignKey:ArticleID;JoinReferences:TagID"`
}

type TagModel struct {
	ID       uint
	Name     string
	Articles []ArticleModel `gorm:"many2many:article_tags;joinForeignKey:TagID;JoinReferences:ArticleID"`
}

type ArticleTagModel struct {
	ArticleID uint `gorm:"primaryKey"` // article_id
	TagID     uint `gorm:"primaryKey"` // tag_id
	CreatedAt time.Time
}

func main() {
	DB.SetupJoinTable(&ArticleModel{}, "Tags", &ArticleTagModel{})
	DB.SetupJoinTable(&TagModel{}, "Articles", &ArticleTagModel{})
	DB.AutoMigrate(&ArticleModel{}, &TagModel{}, &ArticleTagModel{})
	DB.Create(&ArticleModel{
		Title: "计算机科学与技术",
		Tags: []TagModel{
			{Name: "数据结构"},
			{Name: "计算机组成原理"},
		},
	})
}
