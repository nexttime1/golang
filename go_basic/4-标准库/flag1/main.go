package main

import (
	"flag"
	"fmt"
)

func main() {
	// flag.Type() 的使用
	name := flag.String("name", "张三", "姓名")
	age := flag.Int("age", 18, "年龄")
	married := flag.Bool("married", false, "婚否")
	delay := flag.Duration("d", 0, "时间间隔")

	flag.Parse()

	fmt.Println(*name, *age, *married, *delay)
}
