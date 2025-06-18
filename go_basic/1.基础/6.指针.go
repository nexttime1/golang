package main

import "fmt"

func main() {
	var a = 10
	var p = &a                             //每个变量都有地址  p是指针 也有地址  它的值是个地址
	fmt.Printf("a的值为：%v a的地址为%p\n", a, &a) //a的值为：10 a的地址为0xc00008c0a8
	fmt.Printf("p的值为：%v p的地址为%p\n", p, &p) //p的值为：0xc00008c0a8 p的地址为0xc00008e058
	//引用数据类型  必须先分配空间  切片 map 指针 都是  必须先make 或 new
	var q = new(int)
	fmt.Printf("q默认值为%v q地址为%p 指向地址的值为：%v\n", q, &q, *q)
}
