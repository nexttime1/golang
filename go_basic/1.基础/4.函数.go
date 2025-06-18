package main

import (
	"fmt"
	"sort"
)

func sum(x int, y int) int {
	return x + y
}
func sum_some(x int, y ...int) int {
	s := 0
	s += x
	for _, y := range y {
		s += y
	}
	return s
}
func sum1(x, y int) (s int, v int) {
	s = x + y
	v = x - y
	return
}
func SortAsc(dict1 map[string]string) string {
	str := make([]string, 0)
	for k, _ := range dict1 {
		str = append(str, k)
	}
	sort.Strings(str)
	var str1 string
	for _, s := range str {
		str1 += fmt.Sprintf("key => %v value => %v  ", s, dict1[s])
	}
	return str1

}

// 延迟函数
func cul(c string, a, b int) int {
	ret := a + b
	fmt.Println(c, a, b, ret)
	return ret
}

type cla func(int, int) int //自定义类型
type My_Int int

func object(c string) cla {
	switch c {
	case "+":
		return sum
	case "-":
		return func(x, y int) int {
			return x - y
		}
	case "*":
		return func(x, y int) int {
			return x * y
		}
	default:
		return nil
	}

}

// 闭包   不污染全局  变量常驻内存
func pack() func(y int) int {
	var i = 10
	return func(y int) int {
		i += y
		return i
	}
}
func fun3(y int) int {
	x := 5
	defer func(x int) {
		x++
	}(x)
	return x

}

// 抛出异常
func fun4(x, y int) int {

	defer func() error {
		err := recover()
		if err != nil {
			fmt.Println(err)
			return err.(error)
		} else {
			return nil
		}
	}()
	return x / y
}

func main() {
	a := sum(1, 2)
	fmt.Println(a)
	a = sum_some(1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println(a)
	b, c := sum1(2, 2)
	fmt.Println(b)
	fmt.Println(c)
	//首字母升序输出
	dict1 := map[string]string{
		"username": "xtm",
		"age":      "20",
		"sex":      "男",
	}
	str1 := SortAsc(dict1)
	println(str1)
	//自定义类型
	var q cla
	q = sum
	fmt.Println(q(1, 22))
	//匿名函数
	a1 := object("*")
	fmt.Println(a1(1, 2))

	//闭包
	var a2 = pack()
	fmt.Println(a2(10)) //20
	fmt.Println(a2(10)) //30
	fmt.Println(a2(10)) //40
	x := 1
	y := 2
	defer cul("AA", x, cul("A", x, y))
	x = 10
	defer cul("BB", x, cul("B", x, y))
	y = 20
	/*
			注册顺序
			cul("AA", x, cul("A", x, y))
			cul("BB", x, cul("B", x, y))
			执行顺序
			cul("BB", x, cul("B", x, y))
			cul("AA", x, cul("A", x, y))
		执行前所有变量都要知道  所以先算cul("A", x, y)  保证全部都知道  再算cul("B", x, y)
		再算 BB AA
	*/
	fmt.Println(fun3(4)) //5    return 是先赋值  再defer  再返回
	qqq := fun4(10, 5)
	println(qqq)
}
