package time_test

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	now := time.Now()
	year := now.Year()
	month := now.Month()
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()
	fmt.Printf("%02d:%02d:%02d %2d-%2d-%2d\n", year, month, day, hour, minute, second)
	unix := now.Unix() //从1970  之后的时间戳
	fmt.Println(unix)  //秒级别的
	milli := now.UnixMilli()
	fmt.Println(milli) //毫秒级别的
}

func TestFormatTime(t *testing.T) {
	now := time.Now()
	now.Format("2006-01-02 15:04:05")
	fmt.Println(now)

}

func TestDuration(t *testing.T) {
	fiveSecond := time.Second * 5
	now := time.Now()
	time_add := now.Add(fiveSecond)
	fmt.Println(time_add)
	// 做减法  一定要 关注时区
	location, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		panic(err)
	}
	inLocation, _ := time.ParseInLocation("2006-01-02 15:04:05", time_add.Format("2006-01-02 15:04:05"), location)
	fmt.Println(inLocation)
	sub := inLocation.Sub(time.Now())
	fmt.Println(sub)
}

func TestComuper(t *testing.T) {
	now := time.Now()
	now1 := time.Now()
	fmt.Println(now.Equal(now1)) //true
	parse, err := time.Parse("2006-01-02 15:04:05", "2018-08-02 12:04:05")
	if err != nil {
		panic(err)
	}
	fmt.Println(now.After(parse))  //true
	fmt.Println(now.Before(parse)) //false

}

//定时器
/**
*ticker只要定义完成，从此刻开始计时，不需要任何其他的操作，每隔固定时间都会触发。
*timer定时器，是到固定时间后会执行一次
*如果timer定时器要每隔间隔的时间执行，实现ticker的效果，使用 func (t *Timer) Reset(d Duration) bool
 */
func TestTicker(t *testing.T) {
	var wg sync.WaitGroup

	//NewTimer 创建一个 Timer，它会在最少过去时间段 d 后到期，向其自身的 C 字段发送当时的时间
	timer1 := time.NewTimer(2 * time.Second)

	//NewTicker 返回一个新的 Ticker，该 Ticker 包含一个通道字段，并会每隔时间段 d 就向该通道发送当时的时间。它会调
	//整时间间隔或者丢弃 tick 信息以适应反应慢的接收者。如果d <= 0会触发panic。关闭该 Ticker 可
	//以释放相关资源。
	ticker1 := time.NewTicker(2 * time.Second)
	wg.Add(2)

	go func(s *time.Timer) {
		<-s.C
		fmt.Println("每隔两秒 我就出来了  但我只出现一次  除非reset")
		s.Reset(2 * time.Second)
		wg.Done()
	}(timer1)
	go func(s *time.Ticker) {
		<-s.C
		fmt.Println("每隔两秒 我就出来")
		wg.Done()
	}(ticker1)
	wg.Wait()
}
