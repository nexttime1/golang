package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func test_() {
	for i := 0; i < 10; i++ {
		fmt.Printf("test_%d\n", i)
		time.Sleep(time.Millisecond * 50)
	}
	wg.Done()
}

func test_1(x int) {
	for i := (x-1)*3000 + 1; i < x*3000; i++ {
		flag := true
		if i > 1 {
			for j := 2; j < i; j++ {
				if i%j == 0 {
					flag = false
					break
				}

			}
			if flag {
				fmt.Printf("%d是素数\n", i)
			}

		}
	}
	wg.Done()
}

var wg sync.WaitGroup //定义携程

func main() {
	wg.Add(1)  //计数器+1
	go test_() // 并行
	for i := 0; i < 10; i++ {
		fmt.Printf("main_%d\n", i)
		time.Sleep(time.Millisecond * 20)
	}
	//wg.Wait() //等待所有 协程结束
	fmt.Println("done........")

	//获得CPU个数
	fmt.Println(runtime.NumCPU())        //16
	runtime.GOMAXPROCS(runtime.NumCPU()) //设置分配CPU个数

	//用携程去找 0-120000 中的素数
	for i := 1; i <= 4; i++ {
		wg.Add(1)
		go test_1(i)
	}
	wg.Wait()
}
