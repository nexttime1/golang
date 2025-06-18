package byte_ku

import (
	"bytes"
	"fmt"
	"testing"
)

/*
  - `func Split(s, sep []byte) [][]byte`
    Split 以 sep 为分隔符将 s 切分成多个子串，结果不包含分隔符。如果 sep 为空，则将 s 切分成 Unicode 字符列表。
  - `func SplitN(s, sep []byte, n int) [][]byte`
    SplitN 可以指定切分次数 n，超出 n 的部分将不进行切分。
  - `func SplitAfter(s, sep []byte) [][]byte`
    功能同 Split，只不过结果包含分隔符（在各个子串尾部）。
  - `func SplitAfterN(s, sep []byte, n int) [][]byte`
    功能同 SplitN，只不过结果包含分隔符（在各个子串尾部）。
  - `func Fields(s []byte) [][]byte`
    以连续空白为分隔符将 s 切分成多个子串，结果不包含分隔符。
  - `func FieldsFunc(s []byte, f func(rune) bool) [][]byte`
    以符合 f 的字符为分隔符将 s 切分成多个子串，结果不包含分隔符。
  - `func Join(s [][]byte, sep []byte) []byte`
    以 sep 为连接符，将子串列表 s 连接成一个字节串。
  - `func Repeat(b []byte, count int) []byte`
    将子串 b 重复 count 次后返回。
*/
func TestSplit(t *testing.T) {
	b1 := []byte("hello  world!")
	split := bytes.Split(b1, []byte(" "))
	fmt.Printf("%q", split) //["hello" "" "world"]-
	n := bytes.SplitN(b1, []byte(" "), 2)
	fmt.Printf("%q", n)                  //["hello" " world"]--
	fmt.Printf("%q\n", bytes.Fields(b1)) //["hello" "world!"]

	//判断输入的r是否是空格或者叹号（' '或'!'）中的任意一个。如果是，则返回true；否则返回false。
	f := func(r rune) bool { //
		return bytes.ContainsRune([]byte(" !"), r)
	}
	fmt.Printf("%q\n", bytes.FieldsFunc(b1, f)) //["hello" "world"]

}
