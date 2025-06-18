package main

import (
	"fmt"
	"sync"
)

func score(ch chan int) {
	for i := 0; i < 10; i++ {
		x := <-ch
		fmt.Println("取到了一个", x)
	}
	wg1.Done()
}
func load(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	wg1.Done()
}
func group_func(ch chan int) {
	for i := 1; i < 100; i++ {
		ch <- i
	}
	close(ch)
	wg1.Done()
}
func judge_func(GroupCh chan int, judgeCh chan int, expendCh chan bool) {
	for v := range GroupCh { // range 不需要关闭携程
		flag := true
		if v > 1 {
			for j := 2; j < v; j++ {
				if v%j == 0 {
					flag = false
					break
				}

			}
			if flag {
				judgeCh <- v
			}
		}
	}
	expendCh <- true // 表示一个携程完事  16个都完事 才能关闭
	wg1.Done()
}

func printCh_func(ch chan int) {

	for v := range ch {
		fmt.Println("素数为：", v)
	}
	wg1.Done()
}

var wg1 sync.WaitGroup

func main() {
	//管道需要make 且是引用类型
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	a := <-ch
	fmt.Println(a) //队列
	v := <-ch
	fmt.Println(v)

	//使用for循环 写入管道  必须要close   for range 就没啥  建议都关闭
	ch1 := make(chan int, 10)
	for i := 0; i < 10; i++ {
		ch1 <- i
	}
	close(ch1)
	//range 遍历不需要 key
	for v := range ch1 {
		fmt.Println(v)
	}
	//实现生产者和消费者   go自动解决互斥问题
	ch2 := make(chan int, 10)
	wg1.Add(1)
	go score(ch2)
	wg1.Add(1)
	go load(ch2)
	//wg1.Wait()
	//管道实现  16个携程 一块打印素数
	groupCh := make(chan int, 1000)
	judgeCh := make(chan int, 5000)
	expendCh := make(chan bool, 16)
	wg1.Add(1)
	go group_func(groupCh)
	for i := 0; i < 16; i++ {
		wg1.Add(1)
		go judge_func(groupCh, judgeCh, expendCh)
	}

	wg1.Add(1)
	go printCh_func(judgeCh)
	wg1.Add(1)
	go func(expendCh chan bool) {
		for i := 0; i < 16; i++ {
			_ = <-expendCh

		}
		close(expendCh)
		close(judgeCh)
		wg1.Done()
	}(expendCh)
	wg1.Wait()
}
