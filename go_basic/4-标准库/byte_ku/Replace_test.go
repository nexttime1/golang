package byte_ku

import (
	"bytes"
	"fmt"
	"testing"
)

func TestReplace(t *testing.T) {
	s := []byte("hello,worldo")
	old := []byte("o")
	news := []byte("ee")
	// n 代表换几个  0代表不换   负数代表全换
	fmt.Println(string(bytes.Replace(s, old, news, 0)))  //hello,worldo
	fmt.Println(string(bytes.Replace(s, old, news, 1)))  //hellee,worldo
	fmt.Println(string(bytes.Replace(s, old, news, 2)))  //hellee,weerldo
	fmt.Println(string(bytes.Replace(s, old, news, -1))) //hellee,weerldee

	s1 := []byte("你好世界")
	r := bytes.Runes(s1)               //中文一个站3个字节
	fmt.Println("转换前字符串的长度：", len(s1)) //12
	fmt.Println("转换后字符串的长度：", len(r))  //4
}
