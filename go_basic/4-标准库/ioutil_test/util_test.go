package ioutil_test

import (
	"fmt"
	"io"
	"os"
	"testing"
)

//          ioutil  一些函数被移到了io和os包

// 读目录
func TestFunction1(t *testing.T) {
	fi, _ := os.ReadDir(".") // ioutil.ReadDir  弃用
	for _, v := range fi {
		fmt.Printf("v.Name(): %v\n", v.Name())
	}

}

// ReadAll  ioutil   转移到   io
func TestFunction2(t *testing.T) {
	f, _ := os.OpenFile("test.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	all, _ := io.ReadAll(f)
	fmt.Println(string(all))

}

// ReadFile   转移给  os
func TestFunction3(t *testing.T) {
	b, _ := os.ReadFile("test.txt")
	fmt.Printf("string(b): %v\n", string(b))
}
