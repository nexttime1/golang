package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

type Students struct {
	ID    uint    `gorm:"size:10"` //size 的单位是字符 一个字符 1-4哥字节  UTF-8
	Name  string  `gorm:"size:16"`
	Age   int     `gorm:"size:3"`
	email *string `gorm:"size:128"`
	//指针的目的是 可以传空值
}

var mysqlLogger logger.Interface

func init() {
	username := "root"  //账号
	password := "root"  //密码
	host := "127.0.0.1" //数据库地址，可以是Ip或者域名
	port := 3306        //数据库端口
	Dbname := "gorm"    //数据库名
	timeout := "10s"    //连接超时，10秒

	mysqlLogger = logger.Default.LogMode(logger.Info)

	// root:root@tcp(127.0.0.1:3306)/gorm?
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=%s", username, password, host, port, Dbname, timeout)
	//连接MYSQL, 获得DB类型实例，用于后面的数据库读写操作。
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		//为了确保数据一致性，GORM 会在事务里执行写入操作（创建、更新、删除）。如果没有这方面的要求，您可以在初始化时禁用它，这样可以获得60%的性能提升
		SkipDefaultTransaction: true,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "f_",  // 表名前缀
			SingularTable: false, // 单数表名
			NoLowerCase:   false, // 关闭小写转换
		},
		Logger: mysqlLogger,
	})
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}
	// 连接成功
	DB = db
	fmt.Println(db)
}

/*
type 定义字段类型
size 定义字段大小
column 自定义列名
primaryKey 将列定义为主键
unique 将列定义为唯一键
default 定义列的默认值
not null 不可为空
embedded 嵌套字段
embeddedPrefix 嵌套字段前缀
comment 注释
多个标签之前用 ; 连接
*/
type Person struct {
	ID   uint   `gorm:"size:10;primaryKey;column:ID号码"`
	Name string `gorm:"size:5;not null;column:姓名"`
	Age  int    `gorm:"size:3;not null;column:年龄"`
}

//
//func main() {
//	//这是非全局  这个DB就是有日志功能
//	//DB = DB.Session(&gorm.Session{
//	//	Logger: mysqlLogger,
//	//})
//
//	//创建表      AutoMigrate的逻辑是只新增，不删除，不修改（大小会修改）
//	DB.AutoMigrate(&Student{})
//}
