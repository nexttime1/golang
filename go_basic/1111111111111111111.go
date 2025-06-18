package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	//NewTimer 创建一个 Timer，它会在最少过去时间段 d 后到期，向其自身的 C 字段发送当时的时间
	timer1 := time.NewTimer(2 * time.Second)

	//NewTicker 返回一个新的 Ticker，该 Ticker 包含一个通道字段，并会每隔时间段 d 就向该通道发送当时的时间。它会调
	//整时间间隔或者丢弃 tick 信息以适应反应慢的接收者。如果d <= 0会触发panic。关闭该 Ticker 可
	//以释放相关资源。
	ticker1 := time.NewTicker(2 * time.Second)
	wg.Add(2)

	go func(s *time.Timer) {
		defer wg.Done()
		for {
			<-s.C
			fmt.Println("每隔两秒 我就出来了  但我只出现一次  除非reset")
			s.Reset(2 * time.Second)
		}
	}(timer1)
	go func(s *time.Ticker) {
		defer wg.Done()
		for {
			<-s.C
			fmt.Println("每隔两秒 我就出来")
		}

	}(ticker1)
	wg.Wait()
}
