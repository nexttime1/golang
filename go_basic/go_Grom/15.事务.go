package main

import (
	"fmt"
	"gorm.io/gorm"
)

type UserContent struct {
	Name  string
	ID    uint
	Money int
	Age   int
}

func main() {
	DB.AutoMigrate(&UserContent{})
	DB.Create(&UserContent{
		Name:  "张三",
		Money: 100,
		Age:   25,
	})
	DB.Create(&UserContent{
		Name:  "李四",
		Money: 100,
		Age:   24,
	})

	var zhangsan UserContent
	var lisi UserContent
	DB.Find(&zhangsan, 1)
	DB.Find(&lisi, 2)

	//使用事务
	err := DB.Transaction(func(tx *gorm.DB) error {
		err := tx.Model(&zhangsan).Update("money", gorm.Expr("money + 100")).Error
		if err != nil {
			return err
		}
		//err = errors.New("出错了")
		//panic(err)
		tx.Model(&lisi).Update("money", gorm.Expr("money - 100"))
		return nil
	})
	fmt.Println(err)
}
