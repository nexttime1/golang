package main

import (
	"fmt"
	"sort"
)

// 选择排序
func sample(arr []int) {
	min := 0
	for i := 0; i < len(arr); i++ {
		min = i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		t := arr[i]
		arr[i] = arr[min]
		arr[min] = t
	}
}
func sort_sample(arr []int, low int, high int) int {
	tmp := arr[low]
	for low < high {
		for low < high && tmp <= arr[high] {
			high--
		}
		arr[low] = arr[high]
		for low < high && tmp >= arr[low] {
			low++
		}
		arr[high] = arr[low]
	}
	arr[low] = tmp
	return low
}

func quickSort(arr []int, low int, high int) {
	if low < high {
		mid := sort_sample(arr, low, high)
		quickSort(arr, low, mid-1)
		quickSort(arr, mid+1, high)
	}
}
func maopao(arr []int) {
	flag := true
	for i := 0; i < len(arr); i++ {
		flag = false
		for j := len(arr) - 1; j > i; j-- {
			if arr[j] < arr[j-1] {
				flag = true
				t := arr[j]
				arr[j] = arr[j-1]
				arr[j-1] = t
			}
		}
		if flag == false {
			break
		}
	}

}

func main() {
	arr := []int{1, 2, 555, 23, 54, 33, 222, 6, 99, 87, 408}
	arr1 := []float64{1.2, 2.5, 55.22, 23.4, 54.9, 33, 222, 6, 99, 87, 408}
	//sample(arr)
	//quickSort(arr, 0, len(arr)-1)
	//maopao(arr)
	sort.Ints(arr)
	sort.Float64s(arr1)
	fmt.Println(arr1)
	//降序
	sort.Sort(sort.Reverse(sort.IntSlice(arr)))
	fmt.Println(arr)

}
