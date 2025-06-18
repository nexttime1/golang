package byte_ku

import (
	"bytes"
	"fmt"
	"testing"
)

/*
  - `b := bytes.NewBuffer(s []byte)`
    从一个[]byte切片，构造一个Buffer
  - `b := bytes.NewBufferString(s string)`
    从一个string变量，构造一个Buffer
*/
func TestBufer(t *testing.T) {
	buf := bytes.NewBufferString("hello world")
	b := buf.Bytes()
	fmt.Println(string(b))
	var b1 = make([]byte, 6)
	buf.Read(b1)
	fmt.Println(string(b1))
	buf.Write([]byte(" golang!"))
	fmt.Println(buf.String())
	buf.Reset() //清空
	fmt.Println(buf.String())
	buf.Write([]byte("hello world"))
	next := buf.Next(1)       //读出一个字符
	fmt.Println(string(next)) //h
	fmt.Println(buf.String()) // ello world
}
