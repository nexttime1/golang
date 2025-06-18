package main

import (
	"fmt"
	"sync"
)

func text_1() {
	//故意设置错误
	var list []string
	list[0] = "xtm" // 不make 无法赋值
	wgg.Done()
}
func text_2() {
	var list = make([]string, 2)
	list = append(list, "xtm")
	fmt.Println("执行list了")
	wgg.Done()
}

var wgg sync.WaitGroup

func main() {
	//单向管道 只写
	ch1 := make(chan<- int)
	//只读
	ch2 := make(<-chan int)
	println(ch1, ch2)
	//一般用于函数上  某个函数只读  某个函数只写

	ch3 := make(chan int, 10)
	ch4 := make(chan int, 10)
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			ch3 <- i
		} else {
			ch4 <- i
		}
	}
	//select  谁就绪 就执行谁  如果都有 随机  防止饥饿
aaa:
	for {
		select {
		case v := <-ch3:
			fmt.Println(v)
		case v := <-ch4:
			fmt.Println(v)
		default:
			fmt.Println("打印结束！")
			break aaa
		}
	}
	//当 text_1 方法有问题  但程序还要执行  我该怎么办
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("出问题了1111111")
		}
	}()
	wgg.Add(1)
	go text_1()
	wgg.Add(1)
	go text_2()
	wgg.Wait()

}
