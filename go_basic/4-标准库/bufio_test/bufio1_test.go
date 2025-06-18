package bufio_test

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"
	"testing"
)

func TestNewReader(T *testing.T) {
	reader := strings.NewReader("xtm skw")
	reader1 := strings.NewReader("ai_qing")
	b := bufio.NewReader(reader)
	fmt.Println(b.ReadString('\n'))
	b.Reset(reader1)
	fmt.Println(b.ReadString('\n'))

}

func TestReader(t *testing.T) {
	reader := strings.NewReader("xtm skw libai zhaozilong houzi yinggou")
	b := bufio.NewReader(reader)
	for {
		buf := make([]byte, 5)
		n, err := b.Read(buf)
		if n == 0 || err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
		fmt.Println(string(buf[:n])) // 读5个
	}
}

func TestReadLine(t *testing.T) {
	s := strings.NewReader("ABC\nDEF\r\nGHI\r\nGHI")
	br := bufio.NewReader(s)

	w, isPrefix, _ := br.ReadLine()
	fmt.Printf("%q %v\n", w, isPrefix)

	w, isPrefix, _ = br.ReadLine()
	fmt.Printf("%q %v\n", w, isPrefix)

	w, isPrefix, _ = br.ReadLine()
	fmt.Printf("%q %v\n", w, isPrefix)

	w, isPrefix, _ = br.ReadLine()
	fmt.Printf("%q %v\n", w, isPrefix)
}

func TestReadSlice(t *testing.T) {
	s := strings.NewReader("ABC,DEF\r\nGHI\r\nGHI")
	br := bufio.NewReader(s)

	w, _ := br.ReadSlice(',')
	fmt.Printf("%q\n", w)

	w, _ = br.ReadSlice(',')
	fmt.Printf("%q\n", w)

	w, _ = br.ReadSlice(',')
	fmt.Printf("%q\n", w)
}

func TestBufferedReader(t *testing.T) {
	buf1 := bytes.NewBuffer(make([]byte, 0))
	w1 := bufio.NewWriter(buf1)
	w1.WriteString("skw xtm")
	buf2 := bytes.NewBuffer(make([]byte, 0))
	w1.Reset(buf2)
	w1.WriteString("xtm wu di")
	w1.Flush()                 //这个时候才真正写入
	fmt.Println(buf1.String()) //空
	fmt.Println(buf2.String())
}

// 可以 用  ReadFrom  就不用 flush
func TestReadFrom(t *testing.T) {
	buf1 := bytes.NewBuffer(make([]byte, 0))
	writer := bufio.NewWriter(buf1)
	r1 := strings.NewReader("xtm love skw")
	writer.ReadFrom(r1)
	fmt.Println(buf1.String())
}
