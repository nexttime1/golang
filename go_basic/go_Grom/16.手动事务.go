package main

import "gorm.io/gorm"

type PersonInfo struct {
	ID    uint   `gorm:"primary_key"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Money int    `json:"money"`
}

func main() {
	DB.AutoMigrate(&PersonInfo{})
	DB.Create(&PersonInfo{Name: "zhangsan", Age: 18, Money: 100})
	DB.Create(&PersonInfo{Name: "lisi", Age: 18, Money: 100})

	var person1 PersonInfo
	DB.First(&person1, 1)
	var person2 PersonInfo
	DB.First(&person2, 2)

	tx := DB.Begin() //手动开始
	err := DB.Model(&person1).Update("money", gorm.Expr("money - 100")).Error
	if err != nil {
		tx.Rollback() //回滚
		return
	}
	DB.Model(&person2).Update("money", gorm.Expr("money + 100"))
	tx.Commit()

}
