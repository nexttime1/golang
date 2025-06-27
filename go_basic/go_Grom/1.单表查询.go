package main

import (
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
)

func main() {
	DB.AutoMigrate(&Student{}) //创建一个表
	err := DB.Create(&Student{Name: "薛提猛", Age: 22}).Error
	fmt.Println(err) //err = nil 就没事
	email := "1514198964@qq.com"
	s1 := &Student{Name: "宋可为", Age: 21, Email: &email}
	DB.Create(&s1)
	fmt.Println(s1.ID) //自动传回来
	//批量插入
	StudentList := []Student{}
	for i := 0; i < 10; i++ {
		StudentList = append(StudentList, Student{
			Name:  fmt.Sprintf("李白—%d", i),
			Age:   i + 1,
			Email: &email,
		})
	}
	err = DB.Create(&StudentList).Error
	fmt.Println(err)

	//单条查询   先看看日志 输出 sql语句
	DB = DB.Session(&gorm.Session{
		Logger: mysqlLogger,
	})
	var student Student
	DB.Take(&student) //查询第一条
	fmt.Println(student)
	student = Student{}
	//Limit 1
	DB.First(&student)   //  SELECT * FROM `f_students` ORDER BY `f_students`.`id学号` LIMIT 1
	fmt.Println(student) //{1 薛提猛 22 <nil>}
	student = Student{}
	//查询最后一个
	DB.Last(&student)    //SELECT * FROM `f_students` ORDER BY `f_students`.`id学号` DESC LIMIT 1
	fmt.Println(student) //{22 李白—9 10 0xc000296180}

	student = Student{}
	err = DB.Take(&student, 45).Error
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(student)
	//拼接查询 where
	student = Student{}
	DB.Take(&student, "name= ?", "宋可为")
	fmt.Println(student)
	//按照结构体去查询
	student.ID = 2
	DB.Take(&student)
	fmt.Println(student)

	//查询多个
	DataList := []Student{}
	count := DB.Take(&DataList).RowsAffected //有几个
	fmt.Println(count)
	for _, s := range DataList {
		fmt.Println(s)
	}
	data, _ := json.Marshal(DataList)
	fmt.Println(string(data))

	//多表 因为主键去查询
	DB.Find(DataList, 1, 2, 3, 4)         //前四个  主键
	DB.First(DataList, []int{1, 2, 3, 4}) //一样
	DB.Find(DataList, "name in (?)", []string{"薛提猛", "宋可为"})
	fmt.Println(DataList)
}
