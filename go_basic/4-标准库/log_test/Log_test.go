package log_test

import (
	"log"
	"testing"
)

func TestLog(t *testing.T) {
	log.Print("this is a log")
	log.Printf("this is a log: %d", 100) // 格式化输出
	name := "zhangsan"
	age := 20
	log.Println(name, " ", age)
}

func TestPanic(t *testing.T) {
	defer log.Println("aiaiaiaiaiai") //会执行
	log.Print("this is a log")
	log.Panic("this is a error")
	log.Print("结束")
}

func TestFatal(t *testing.T) {
	defer log.Println("aiaiaiaiaiai") // 不会执行  fatal 直接 exist 退出程序
	log.Print("this is a log")
	log.Fatal("this is a error")
	log.Print("结束")

}

func TestPre(t *testing.T) {
	log.SetPrefix("Mylog   ")
	log.Print("this is a log")

}
