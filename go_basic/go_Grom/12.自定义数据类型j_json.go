package main

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

type Info struct { //必须要实现两个方法
	Status string `json:"status"`
	Addr   string `json:"addr"`
	Age    int    `json:"age"`
}

// Scan 从数据库中读取出来
func (i *Info) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	info := Info{}
	err := json.Unmarshal(bytes, &info)
	*i = info
	return err
}

// Value 存入数据库
func (i Info) Value() (driver.Value, error) {
	fmt.Printf("存入数据库之前：%v  %#T\n", i, i)
	return json.Marshal(i)
}

type AuthModel struct {
	ID   uint
	Name string
	Info Info `gorm:"type:string"`
}

func main() {
	//DB.AutoMigrate(&AuthModel{})
	//DB.Create(&AuthModel{
	//	Name: "薛提猛",
	//	Info: Info{
	//		Status: "良好",
	//		Addr:   "山东省",
	//	},
	//})
	var authModel AuthModel
	DB.Find(&authModel, 1)
	fmt.Println(authModel)
}
