package main

import "fmt"

func main() {
	str := "localhost"
	number := 8080
	//拼接作用
	addr := fmt.Sprintf("%v:%d", str, number)
	fmt.Println(addr)

	//打印错误更方便
	err := fmt.Errorf("你给出的%s套接字不正确", addr)
	fmt.Println(err)

}
