package main

import (
	"fmt"
	"gorm.io/gorm"
	"time"
)

type Label struct {
	ID       uint `gorm:"primaryKey"`
	Name     string
	Passages []Passage `gorm:"many2many:passage_labels"`
}

// Essay Tag 多对多的关系  这种关系需要一个中间表来关联两个模型，而这个中间表的名称就是 essay_tags
type Passage struct {
	ID     uint `gorm:"primaryKey"`
	Title  string
	Labels []Label `gorm:"many2many:passage_labels;"` //标签来指明这是一个多对多关系
}
type PassageLabel struct {
	LabelID   uint      `gorm:"primaryKey"`
	PassageID uint      `gorm:"primaryKey"`
	CreateAt  time.Time `json:"create_at"`
}

func (at *PassageLabel) BeforeCreate(db *gorm.DB) (err error) {
	at.CreateAt = time.Now()
	return nil
}
func main() {
	//放开这两个 他就知道 我的自定义关系表  就去调用钩子函数
	DB.SetupJoinTable(&Passage{}, "Labels", &PassageLabel{}) //有了这句话  不需要 many2many
	DB.SetupJoinTable(&Label{}, "Passages", &PassageLabel{}) //双向注册
	//DB.AutoMigrate(&Passage{}, &Label{}, &PassageLabel{})

	//添加
	DB.Create(&Passage{
		Title: "计算机科学与技术",
		Labels: []Label{
			{Name: "Python"},
			{Name: "Go"},
			{Name: "Java"},
		},
	})
	DB.Create(&Label{
		Name: "大数据",
	})
	DB.Create(&Passage{
		Title: "软件工程",
	})
	//已知 Passage  添加title  和关系
	var p1 []Passage
	DB.Find(&p1, "title = ?", "软件工程")
	DB.Create(&Label{
		Name:     "数据挖掘",
		Passages: p1,
	})
	// 将原来的 Label  替换
	var l1 []Label
	DB.Find(&l1, []int{4, 7})
	fmt.Println(l1)
	var p2 []Passage
	DB.Find(&p2, "title = ?", "软件工程")
	DB.Model(&p2).Association("Labels").Replace(&l1)

}
