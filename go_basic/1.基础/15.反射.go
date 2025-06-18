package main

import (
	"fmt"
	"reflect"
)

type MyInt int
type Person1 struct {
	name string
	age  int
}

// reflect
func typeInfo(x interface{}) {
	v := reflect.TypeOf(x)             // 比如 person   v 为main.person v.name 为person   v.Kind 为 struct
	fmt.Println(v, v.Name(), v.Kind()) //v 为  什么类型   v.Name 为名字是什么  v.Kind 是底层是什么

}

func value_of(x interface{}) {
	v := reflect.ValueOf(x)
	switch v.Kind() {
	case reflect.Int:
		fmt.Println("integer为", v.Int())
	case reflect.String:
		fmt.Println("string 为", v.String())
	default:
		fmt.Println("未识别到")
	}

}
func Set_Value(x interface{}) {
	v := reflect.ValueOf(x)
	switch v.Elem().Kind() { //Elem() 专门应对指针变量  v.Kind 是ptr
	case reflect.Int:
		v.Elem().SetInt(520)
	case reflect.String:
		v.Elem().SetString("MOM")
	default:
		fmt.Println("未识别到")
	}

}

func main() {
	a := 10
	b := 12.5
	c := true
	d := "xtm"
	typeInfo(a)
	typeInfo(b)
	typeInfo(c)
	typeInfo(d)
	var e MyInt = 222
	f := Person1{name: "xtm", age: 18}
	typeInfo(e)
	typeInfo(f)
	g := [4]int{1, 2, 3, 4}
	typeInfo(g)            //数组
	h := []int{1, 2, 3, 4} //切片
	typeInfo(h)

	//reflect valueOf
	value_of(a)

	//想要修改咋办
	Set_Value(&a)
	fmt.Println(a)
	Set_Value(&d)
	fmt.Println(d)
}
