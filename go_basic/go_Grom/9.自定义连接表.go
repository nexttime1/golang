package main

import (
	"fmt"
	"time"
)

//不双向引用的案例

type Paper struct {
	ID      uint
	Title   string
	Tallies []Tally `gorm:"many2many:paper_tallies;"`
}

type Tally struct {
	ID   uint
	Name string
}

type PaperTally struct {
	PaperID   uint `gorm:"primaryKey"`
	TallyID   uint `gorm:"primaryKey"`
	CreatedAt time.Time
}

func main() {
	// 设置Article的Tags表为ArticleTag
	DB.SetupJoinTable(&Paper{}, "Tallies", &PaperTally{})
	// 如果tag要反向应用Article，那么也得加上
	//DB.SetupJoinTable(&Tally{}, "Papers", &PaperTally{})
	err := DB.AutoMigrate(&Paper{}, &Tally{}, &PaperTally{})
	fmt.Println(err)
	DB.Create(&Paper{
		Title: "flask零基础入门",
		Tallies: []Tally{
			{Name: "python"},
			{Name: "后端"},
			{Name: "web"},
		},
	})
	//有了Tallies  添加Paper
	var p1 Paper
	DB.Find(&p1, "title = ?", "GO语言")
	var t1 []Tally
	DB.Find(&t1, []int{4, 5})
	DB.Model(&p1).Association("Tallies").Append(t1)

	//替换
	var p2 Paper
	DB.Find(&p2, "title = ?", "GO语言")
	var t2 []Tally
	DB.Find(&t2, []int{2, 3})
	DB.Model(&p2).Association("Tallies").Replace(t2)
}
