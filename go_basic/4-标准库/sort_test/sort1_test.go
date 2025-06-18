package sort_test

import (
	"fmt"
	"sort"
	"testing"
)

type NewInt []int

func TestSort1(T *testing.T) {

	a := []int{1, 4, 6, 33, 1, 2, 408, 9}
	sort.Ints(a)
	fmt.Println(a)
	//先 转成  IntSlice类型   然后reverse  在sort
	sort.Sort(sort.Reverse(sort.IntSlice(a)))
	fmt.Println(a)

	//search
	index := sort.SearchInts([]int{1, 2, 3, 4}, 4) //必须排序好的
	fmt.Println(index)
}

// string 排序
func TestSort2(T *testing.T) {
	ls := sort.StringSlice{ //变成StringSlice 类型
		"100",
		"42",
		"41",
		"3",
		"2",
	}
	fmt.Println(ls)
	sort.Strings(ls)
	fmt.Println(ls)
}
