package main

import (
	"fmt"
	"time"
)

func main() {
	data := time.Now()
	//fmt.Println(data)  2025-05-18 16:09:00.5984942 +0800 CST m=+0.000000001
	year := data.Year()
	month := data.Month()
	day := data.Day()
	hour := data.Hour()
	minute := data.Minute()
	second := data.Second()
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)

	str := data.Format("2006-01-02 15:04:05") //24小时制
	fmt.Println(str)                          //	2025-05-18 16:11:59
	str = data.Format("2006/01/02 03:04:05")  //12小时制
	fmt.Println(str)                          //2025/05/18 04:11:59
	unix := data.Unix()                       //时间戳  毫秒
	fmt.Println(unix)
	str = "2025-03-02 04:05:12"
	tmp := "2006-01-02 15:04:05"
	location, _ := time.ParseInLocation(tmp, str, time.Local)
	fmt.Println(location) //2025-03-02 04:05:12 +0800 CST
	unix = location.Unix()
	fmt.Println(unix)

	//定时器
	ticker := time.NewTicker(time.Second) // 每一秒执行一次
	n := 5
	for t := range ticker.C {
		if n == 0 {
			ticker.Stop()
			break
		}
		fmt.Println(t)
		n--
	}

}
