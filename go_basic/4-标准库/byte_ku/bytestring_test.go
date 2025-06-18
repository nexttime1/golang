package byte_ku

import (
	"bytes"
	"fmt"
	"testing"
)

/*
  - `func HasPrefix(s, prefix []byte) bool`
    判断 s 是否有前缀 prefix
  - `func HasSuffix(s, suffix []byte) bool`
    判断 s 是否有后缀 suffix
  - `func Contains(b, subslice []byte) bool`
    判断 b 中是否包含子串 subslice
  - `func ContainsRune(b []byte, r rune) bool`
    判断 b 中是否包含子串 字符 r
  - `func ContainsAny(b []byte, chars string) bool`
    判断 b 中是否包含 chars 中的任何一个字符
  - `func Index(s, sep []byte) int`
    查找子串 sep在 s 中第一次出现的位置，找不到则返回 -1
  - `func IndexByte(s []byte, c byte) int`
    查找子串 字节 c在 s 中第一次出现的位置，找不到则返回 -1
  - `func IndexRune(s []byte, r rune) int`
    查找子串字符 r在 s 中第一次出现的位置，找不到则返回 -1
  - `func IndexAny(s []byte, chars string) int`
    查找 chars 中的任何一个字符在 s 中第一次出现的位置，找不到则返回 -1。
  - `func IndexFunc(s []byte, f func(r rune) bool) int`
    查找符合 f 的字符在 s 中第一次出现的位置，找不到则返回 -1。
  - `func LastIndex(s, sep []byte) int`
    功能同上，只不过查找最后一次出现的位置。
  - `func LastIndexByte(s []byte, c byte) int`
    功能同上，只不过查找最后一次出现的位置。
  - `func LastIndexAny(s []byte, chars string) int`
    功能同上，只不过查找最后一次出现的位置。
  - `func LastIndexFunc(s []byte, f func(r rune) bool) int`
    功能同上，只不过查找最后一次出现的位置。
  - `func Count(s, sep []byte) int`
    获取 sep 在 s 中出现的次数（sep 不能重叠）
*/
func TestString(t *testing.T) {
	b := []byte("hello golang") //字符串强转为byte切片
	sublice1 := []byte("hello")
	sublice2 := []byte("Hello")
	fmt.Println(bytes.Contains(b, sublice1)) //true
	fmt.Println(bytes.Contains(b, sublice2)) //false

	s := []byte("hellooooooooo")
	sep1 := []byte("h")
	sep2 := []byte("l")
	sep3 := []byte("o")
	fmt.Println(bytes.Count(s, sep1))        //1
	fmt.Println(bytes.Count(s, sep2))        //2
	fmt.Println(bytes.Count(s, sep3))        //9
	fmt.Println(bytes.Index(s, []byte("o"))) //4
}
