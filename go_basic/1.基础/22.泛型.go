package main

import (
	"fmt"
)

func add_111[T int | float64](x, y T) T {
	return x + y
}

// 自定义切片
type myList[T int | int32 | int64 | float64] []T

// 自定义map
type MyMap[K comparable, V int | int32 | int64 | float64 | string] map[K]V

// 泛型结构体
type MyStruct[T int | string] struct {
	Name T
	Age  int
}

type Conpany[T int | float64, S []int | []string] struct {
	NO  T
	abc S
}

func main() {
	var a int = 14
	var b int = 15
	fmt.Println(add_111(a, b))
	var ccc1 = 12.5
	var d = 11.2
	fmt.Println(add_111(d, ccc1))
	m := myList[int]{1, 2, 3}
	fmt.Println(m[0])
	fmt.Println(1.0 + int64(2))
	m2 := MyMap[int, string]{1: "xtm", 2: "skw"}
	fmt.Println(m2)
	m3 := make(MyMap[float64, int], 0)
	m3[2.2] = 10
	fmt.Println(m3)
	c := Conpany[int, []int]{
		NO:  1,
		abc: []int{1, 2, 3, 4},
	}
	fmt.Println(c)
}
