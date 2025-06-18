package main

import "fmt"

// 实现多个接口
type Get interface {
	GetName() string
}
type Set interface {
	SetName(string)
}

type Dog1 struct {
	Name string
}

func (d *Dog1) GetName() string {
	return d.Name
}
func (d *Dog1) SetName(name string) {
	d.Name = name
}

// 接口的嵌套
type Animaler interface {
	Get
	Set
}

func main() {
	var d1 = Dog1{"xtm"}
	var D1 Get = &d1
	var D2 Set = &d1
	D2.SetName("SKW")
	fmt.Println(D1.GetName())
	var animal Animaler = &d1
	animal.SetName("薛提猛")
	fmt.Println(animal.GetName())

}
