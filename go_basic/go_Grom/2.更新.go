package main

import (
	"fmt"
	"gorm.io/gorm"
)

// 钩子函数
func (user *Student) BeforeCreate(tx *gorm.DB) (err error) {
	email := fmt.Sprintf("%s@qq.com", user.Name)
	user.Email = &email
	return nil
}
func main() {
	DB = DB.Session(&gorm.Session{
		Logger: mysqlLogger,
	})
	//save  全部更新
	var s1 Student
	DB.Find(&s1, 6)
	fmt.Println(s1)
	s1.Name = "橙白"
	//指定的话
	//DB.Select("name").Save(&s1)
	DB.Save(&s1) //UPDATE `f_students` SET `姓名`='橙白',`年龄`=21,`邮件`='1514198964@qq.com' WHERE `id学号` = 6

	//更新多个列
	dataList := []Student{}
	DB.Find(&dataList, []int{6, 13, 8}).Update("age", 22) //主键为6 8 13 的
	//更新多个列 并且多个值
	DB.Find(&dataList, []int{2, 3, 4}).Updates(Student{Name: "一样名字", Age: 18}) //结构体的方式
	DB.Find(&dataList, []int{2, 3, 4}).Updates(map[string]interface{}{"age": 22, "email": "123@qq.com"})

	//删除   先查在删除
	var s2 Student
	DB.First(&s2, 1)
	DB.Delete(&s2)
}
