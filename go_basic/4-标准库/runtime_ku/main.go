package runtime_ku

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(2)
	go func(s string) {
		for i := 0; i < 2; i++ {
			fmt.Println(s)
			runtime.Gosched() //让出CPU
		}
		wg.Done()
	}("world")
	go func(s string) {
		for i := 0; i < 2; i++ {
			fmt.Println(s)
			runtime.Gosched()
		}
		wg.Done()
	}("xtm skw")
	wg.Wait()
}
