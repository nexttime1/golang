package main

import "fmt"

// 接口不需要显示的写出  只要有那些方法  自动就实现接口
type USB interface {
	start()
	end()
}

// 测试 是否显式的写出
func Computer(usb USB) {
	_, ok := usb.(Phone)
	if ok {
		usb.start()
	} else {
		usb.end()
	}
}

// 手机  和  照相机 实现接口
type Phone struct {
	Name string
}

func (phone Phone) start() {
	fmt.Println(phone.Name, "启动！")
}
func (phone Phone) end() {
	fmt.Println(phone.Name, "熄火！")
}

// 自动 var u1 USB = pi Phone
type photo struct {
}

func (photo photo) start() {
	fmt.Println("照相机启动！！！")
}
func (photo photo) end() {
	fmt.Println("照相机结束！！！")
}

// 空接口 可以 是任意类型
type A interface {
}

func main() {
	var p1 = Phone{Name: "小米"}
	var ph = photo{}
	var u1 USB = p1 //多态的形式 去显示的实现接口
	u1.start()
	u1.end()
	ph.start() //自动的实现接口
	ph.end()

	var a A //任意格式
	a = "1"
	a = 123
	a = true
	fmt.Println(a)
	var list = []interface{}{1, 2, "3", true}
	fmt.Println(list)
	var map_ = make(map[string]interface{})
	map_["name"] = "xtm"
	map_["age"] = 18
	fmt.Println(map_)
	Computer(ph)

}
