package byte_ku

import (
	"bytes"
	"fmt"
	"testing"
)

/*
* `func Trim(s []byte, cutset string) []byte`
  去掉 s 两边包含在 cutset 中的字符（返回 s 的切片
* `func TrimLeft(s []byte, cutset string) []byte`
  去掉 s 左边包含在 cutset 中的字符（返回 s 的切片）
* `func TrimRight(s []byte, cutset string) []byte`
  去掉 s 右边包含在 cutset 中的字符（返回 s 的切片）
* `func TrimFunc(s []byte, f func(r rune) bool) []byte`
  去掉 s 两边符合 f函数====返回值是true还是false 要求的字符（返回 s 的切片）
*/

func TestFunc(t *testing.T) {
	//b1 := []byte("hello world. !")
	b2 := []byte("! hello golang. ")
	b3 := []byte("! he!llo xtm !")
	trim1 := bytes.Trim(b3, "!")     // he!llo xtm  只是去掉两边
	trim2 := bytes.TrimLeft(b3, "!") //he!llo xtm !
	fmt.Println(string(trim1), string(trim2))
	trim3 := bytes.TrimRight(b2, "!") //! hello golang.
	fmt.Println(string(trim3))

	//
	trimFunc := bytes.TrimFunc(b2, func(r rune) bool {
		return bytes.ContainsRune([]byte("!！. "), r)
	})
	fmt.Println(string(trimFunc))

}
