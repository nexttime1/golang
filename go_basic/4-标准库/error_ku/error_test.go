package error_ku

import (
	"fmt"
	"testing"
)

/*
自定义error  只要实现这个接口就行

	type error interface {
	    Error() string
	}
*/
type MyError struct {
	code int
	msg  string
}

func (e *MyError) Error() string {
	e.code = 500
	e.msg = "字符串不能为空"
	return e.msg
}

func Check(s string) (int, bool) {
	if s == "" {
		err := MyError{}
		fmt.Printf("错误是：%s 错误码为：%d\n", err.Error(), err.code)
		return err.code, false
	}
	return 200, true
}

func TestError(t *testing.T) {
	s := ""
	check, e := Check(s)
	fmt.Println(check)
	fmt.Println(e)

}
