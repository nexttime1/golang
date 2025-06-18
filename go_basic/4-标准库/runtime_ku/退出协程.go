package runtime_ku

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg1 sync.WaitGroup
	wg1.Add(1)
	go func() {
		defer wg1.Done()
		defer fmt.Println("A.defer")
		func() {
			defer fmt.Println("B.defer")
			// 结束协程
			runtime.Goexit()
			defer fmt.Println("C.defer")
			fmt.Println("B")
		}()
		fmt.Println("A")
	}()

	wg1.Wait() // 主goroutine等待子goroutine结束，主在结束
}

//只打印  fmt.Println("B.defer")        和   defer fmt.Println("A.defer")
//因为  携程结束后面的不执行了
