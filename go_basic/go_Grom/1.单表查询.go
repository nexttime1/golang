package main

import "fmt"

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

}
