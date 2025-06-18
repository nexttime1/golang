package main

import (
	"fmt"
	"reflect"
)

type UU struct {
	Name  string `json:"name" form:"username"`
	Age   int    `json:"age"`
	Hobby string `json:"hobby"`
}

func (u UU) Wang() {
	fmt.Println(u.Name, "正在旺旺旺")
}

func (u *UU) SetInfo(name string, age int, hobby string) {
	u.Name = name
	u.Age = age
	u.Hobby = hobby
}

func Struct_Value(x interface{}) {
	v := reflect.TypeOf(x)
	//value of
	val := reflect.ValueOf(x)
	if v.Kind() != reflect.Struct && v.Elem().Kind() != reflect.Struct {
		fmt.Println("不是结构体")
		return
	}
	if v.Elem().Kind() == reflect.Struct {
		v = v.Elem()
		val = val.Elem()
	}
	//通过类型变量中的 Field 获得 struct的内容
	field0 := v.Field(0)
	fmt.Printf("%#v\n", field0)
	fmt.Printf("%v\n", field0.Name)
	fmt.Printf("%v\n", field0.Type)
	fmt.Printf("%v\n", field0.Tag.Get("json"))

	//通过类型里的 FieldByName
	field1, _ := v.FieldByName("Age")
	fmt.Printf("%#v\n", field1)
	fmt.Printf("%v\n", field1.Name)
	fmt.Printf("%v\n", field1.Type)
	fmt.Printf("%v\n", field1.Tag.Get("json"))
	//求个数
	num := v.NumField()
	fmt.Printf("有%v个字段\n", num)

	name := val.FieldByName("Name")
	fmt.Println(name)
	//便利所有字段
	for i := 0; i < num; i++ {
		fmt.Printf("字段名称：%v 字段类型：%v 字段值：%v Tag：%v  \n", v.Field(i).Name, v.Field(i).Type, val.Field(i), v.Field(i).Tag.Get("json"))
	}
}

// 反射方法
func reflect_Method(x interface{}) {
	v := reflect.TypeOf(x)
	val := reflect.ValueOf(x)
	if v.Kind() != reflect.Struct && v.Elem().Kind() != reflect.Struct {
		fmt.Println("不符合")
		return
	}
	if v.Elem().Kind() == reflect.Struct {
		v = v.Elem()
		//val = val.Elem()   不能加  要不识别不到  指针的方法
	}
	Method0 := v.Method(0) //按 ASCII 编排
	fmt.Printf("%v\n", Method0.Name)
	fmt.Printf("%v\n", Method0.Type)
	//用名字 寻找  并执行  要用  value 	//调用方法需要传参
	var params []reflect.Value
	params = append(params, reflect.ValueOf("next_time"))
	params = append(params, reflect.ValueOf(22))
	params = append(params, reflect.ValueOf("斗破苍穹"))
	mothod := val.MethodByName("SetInfo")
	mothod.Call(params)

}

// 反射 修改熟属性值
func Sett_Value(x interface{}) {
	v := reflect.TypeOf(x)
	val := reflect.ValueOf(x)
	if v.Kind() != reflect.Ptr {
		fmt.Println("请传入指针型的 结构体")
		return
	} else if v.Elem().Kind() != reflect.Struct {
		fmt.Println("请传入结构体")
		return
	}
	name := val.Elem().FieldByName("Name")
	name.SetString("XTM")
	age := val.Elem().FieldByName("Age")
	age.SetInt(20)
}
func main() {
	u := UU{
		Name:  "xtm",
		Age:   18,
		Hobby: "打游戏",
	}
	u.SetInfo("www", 18, "<UNK>")
	fmt.Println(u)
	Struct_Value(&u)
	v := reflect.TypeOf(u)
	method := v.NumMethod()
	fmt.Println(method)
	reflect_Method(&u)
	fmt.Println(u)
	Sett_Value(&u)
	fmt.Println(u)
}
