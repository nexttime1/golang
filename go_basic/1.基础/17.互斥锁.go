package main

import (
	"fmt"
	"sync"
	"time"
)

var mutex1 sync.Mutex   // 生产者 消费者
var mutex2 sync.RWMutex //读者 写者 问题
var wagg1 sync.WaitGroup
var count = 0

func t1() {
	mutex1.Lock()
	count++
	println(count)
	mutex1.Unlock()
	wagg1.Done()
}

func reader() {
	mutex2.RLock()
	time.Sleep(time.Second)
	fmt.Println("reader 在读了....")
	mutex2.RUnlock()
	wagg1.Done()
}
func writer() {
	mutex2.Lock()
	time.Sleep(time.Second)
	fmt.Println("writer 在写小作文了....")
	mutex2.Unlock()
	wagg1.Done()
}

func main() {
	//测试  正常 互斥
	for i := 0; i < 10; i++ {
		wagg1.Add(1)
		go t1()
	}
	time.Sleep(time.Second * 5)
	for i := 0; i < 10; i++ {
		wagg1.Add(1)
		go reader()
		wagg1.Add(1)
		go writer()
	}
	wagg1.Wait()

}
