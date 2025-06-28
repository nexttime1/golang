package main

import (
	"fmt"
	"gorm.io/gorm"
)

func main() {
	DB = DB.Session(&gorm.Session{
		Logger: mysqlLogger,
	})
	var ss1 []Student
	//排序
	DB.Order("age desc").Find(&ss1)
	DB.Order("age asc").Find(&ss1)
	fmt.Println(ss1)

	//分页  每页2分  第0页
	DB.Limit(2).Offset(0).Find(&ss1)
	fmt.Println(ss1)

	//公式去解决
	limit := 2
	page := 4
	DB.Limit(limit).Offset((page - 1) * limit).Find(&ss1)

	//去重  两种
	DB.Model(Student{}).Select("年龄").Distinct().Scan(&ss1)
	fmt.Println(ss1)
	DB.Select("distinct 年龄").Find(&ss1)
	fmt.Println("11111:::::\n", ss1)

	//分组查询
	type Group struct {
		Count    uint `gorm:"column:count"`
		age      int  `gorm:"column:年龄"`
		NameList string
	}
	var groups []Group
	DB.Model(Student{}).Select("group_concat(姓名) as NameList,年龄,count(*) as count").Group("年龄").Find(&groups)
	fmt.Println(groups)

	//原生sql
	DB.Raw("select * from students where age= ? ", 18).Scan(&ss1)
	fmt.Println(ss1)

	//子查询   年龄大于平均年龄的
	DB.Where("age > (?)", DB.Model(Student{}).Select("avg(年龄)")).Find(&ss1)
	fmt.Println(ss1)

	//map的方式  执行SQL 语句
	DB.Where("name = ? and age = ?", map[string]any{"name": "xtm", "age": 18}).Find(&ss1)

	//查询引用  Scope
	DB.Model(Student{}).Scopes(ABC123).Find(&ss1) //Scope 里面函数 必须形参和返回 必须是 *gorm.DB

}
func ABC123(db *gorm.DB) *gorm.DB {
	return db.Where("age > ?", 18)
}
