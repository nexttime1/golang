package io_ku

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"
)

func TestReader(t *testing.T) {
	open, err := os.Open("./test.txt")
	if err != nil {
		panic(err)
	}
	defer open.Close()
	var body []byte
	open.Seek(1, 0)
	for {
		buf := make([]byte, 3)
		n, err := open.Read(buf)
		if err == io.EOF {
			fmt.Println("读完成")
			break
		} else if err != nil {
			panic(err)
		}
		body = append(body, buf[:n]...)
	}
	fmt.Println(string(body))
}
func TestCopy(t *testing.T) {
	open, err := os.Open("./test.txt")
	file, err := os.OpenFile("./test1.txt", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	defer open.Close()
	io.Copy(file, open)

}
func TestLimitReader(t *testing.T) {
	reader := bytes.NewReader([]byte("hello world"))
	limitReader := io.LimitReader(reader, int64(5))
	buf := make([]byte, 10)
	fmt.Println(limitReader.Read(buf)) //  最多读5个
	fmt.Println(string(buf[:5]))
}

func TestMultReader(t *testing.T) {
	r1 := bytes.NewReader([]byte("hello world  "))
	r2 := bytes.NewReader([]byte("hello xtm  "))
	r3 := bytes.NewReader([]byte("hello skw"))
	reader := io.MultiReader(r1, r2, r3)
	io.Copy(os.Stdout, reader) //全部读出

}
func TestMultWriter(t *testing.T) {
	reader := bytes.NewReader([]byte("hello world"))
	var buf1, buf2 bytes.Buffer
	writer := io.MultiWriter(&buf1, &buf2)
	io.Copy(writer, reader)
	fmt.Println(string(buf1.String()))
	fmt.Println(string(buf2.String()))

}
