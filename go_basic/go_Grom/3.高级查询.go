package main

import (
	"fmt"
	"gorm.io/gorm"
)

func main() {
	var studentList []Student
	studentList = []Student{
		{ID: 1, Name: "李元芳", Age: 32, Email: PtrString("lyf@yf.com")},
		{ID: 2, Name: "张武", Age: 18, Email: PtrString("zhangwu@lly.cn")},
		{ID: 3, Name: "枫枫", Age: 23, Email: PtrString("ff@yahoo.com")},
		{ID: 4, Name: "刘大", Age: 54, Email: PtrString("liuda@qq.com")},
		{ID: 6, Name: "李琦", Age: 14, Email: PtrString("liqi@lly.cn")},
		{ID: 7, Name: "晓梅", Age: 25, Email: PtrString("xiaomeo@sl.com")},
		{ID: 8, Name: "如燕", Age: 26, Email: PtrString("ruyan@yf.com")},
		{ID: 9, Name: "魔灵", Age: 21, Email: PtrString("moling@sl.com")},
	}
	DB.Create(&studentList) //Inset into
	DB.Session(&gorm.Session{
		Logger: mysqlLogger,
	})
	var s3 Student
	//姓名= 薛提猛
	Count1 := DB.Find(&s3, "name = ?", "薛提猛").RowsAffected
	DB.Where("name = ?", "薛提猛").Find(&s3)
	fmt.Println(s3, Count1)
	var StudentList []Student
	DB.Where("name in ?", []string{"薛提猛", "宋可为"}).Find(&StudentList)
	//模糊匹配
	DB.Where("name like ?", "李%").Find(&StudentList)
	// qq邮箱 + age > 23
	DB.Where("age > 23 and email like ", "%@qq.com").Find(&StudentList)
	// qq 邮箱 或者 age >20
	DB.Where("age > ? or email like", 23, "%@qq.com").Find(&StudentList)
	//或者
	DB.Where("age > ?", 23).Or("email like", "%@qq.com").Find(&StudentList)
	//结构体查询
	DB.Where(&Student{
		Name: "薛提猛",
		Age:  0, //0值 它会忽略
	}).Find(&StudentList)
	fmt.Println(StudentList, Count1)

	//select 字段可以投影指定列
	DB.Select("name", "age").Find(&studentList)
	//扫描   Model 指定表的意思  就是按结构体去找表
	DB.Model(&Student{}).Select("name", "age").Where("name = ?", "薛提猛").Scan(&studentList)

}
func PtrString(email string) *string {
	return &email
}
