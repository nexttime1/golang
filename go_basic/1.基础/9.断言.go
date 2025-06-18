package main

import (
	"fmt"
)

func info(x interface{}) {
	switch x.(type) {
	case string:
		fmt.Println("string...")
	case int:
		fmt.Println("int...")
	case float64:
		fmt.Println("float64...")
	default:
		fmt.Println("bool")
	}

}

// 如果把切片给到空接口  无法使用其下标
type User_ struct {
	Name  string
	Age   int
	hobby interface{}
	Addre interface{}
}
type Address_ struct {
	street string
	city   string
}

func main() {
	info("<UNK>")
	info(1)

	var u1 = User_{
		Name:  "xtm",
		Age:   18,
		hobby: []string{"吃饭", "玩游戏"},
		Addre: Address_{"<大街>", "<城市>"},
	}
	//fmt.Println(u1.hobby[0])  //错误格式
	list, ok := u1.hobby.([]string)
	if ok {
		fmt.Println(list[0]) //这样才对
	}
	addr, ok := u1.Addre.(Address_)
	if ok {
		fmt.Println(addr.street)
	}
}
