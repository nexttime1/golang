package byte_ku

import (
	"bytes"
	"fmt"
	"testing"
)

func TestByte(t *testing.T) {
	//b := []byte("薛提猛和宋可为")
	b := []byte("hello")

	//转大写
	upper := bytes.ToUpper(b)
	fmt.Println(string(upper))
	fmt.Println(string(b))

	//小写
	c := []byte("LIBAI")
	lower := bytes.ToLower(c)
	fmt.Println(string(lower))

	//比较  compared
	a1 := []byte("ABC")
	a2 := []byte("ABc")
	compare := bytes.Compare(a1, a2) //返回int  大于 为 1  相等为 0  小于  为-1
	fmt.Println(compare)

	//相等
	equal := bytes.Equal(a1, a2) //返回bool
	fmt.Println(equal)

	//忽略大小写  比较
	fold := bytes.EqualFold(a1, a2)
	fmt.Println(fold)
}
