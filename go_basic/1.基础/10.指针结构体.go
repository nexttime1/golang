package main

import "fmt"

type Anminal interface {
	GetName() string
	SetName(name string)
}
type Dog_new struct {
	name string
}

func (d *Dog_new) GetName() string {
	return d.name
}
func (d *Dog_new) SetName(name string) {
	d.name = name
}

type Cat struct {
	name string
}

func (d *Cat) GetName() string {

	return d.name
}
func (d *Cat) SetName(name string) {
	d.name = name
}

func main() {
	var c1 = &Cat{name: "xxx"}
	var d1 = &Dog_new{name: "ddd"}
	c1.SetName("宋可为")
	d1.SetName("薛提猛")
	fmt.Println(c1.GetName())
	fmt.Println(d1.GetName())
}
