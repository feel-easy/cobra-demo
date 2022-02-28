package demo

import (
	"fmt"
	"unsafe"
)

func SliceDemo() {
	var arr = [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	var arr0 [5]int = [5]int{1, 2, 3}
	var s2 = make([]byte, 5, 6)
	s2 = append(s2, 'a')
	s2 = append(s2, 'b')
	s2 = append(s2, 'c')
	arr1 := arr0[:]
	arr1 = append(arr1, 100, 101, 102, 103, 104, 105, 106, 107, 108, 109, 200)
	fmt.Println(arr, arr0, s2, arr1, cap(arr1))
	var ptr unsafe.Pointer
	var s1 = struct {
		addr uintptr
		len  int
		cap  int
	}{uintptr(ptr), 10, 10}
	s := *(*[]byte)(unsafe.Pointer(&s1))
	s3 := arr[:3]
	s3 = append(s3, 100)
	fmt.Println(len(s), s3, arr)
	arr3 := []int{1, 9, 10, 30, 2, 5, 45, 8, 63, 234, 12}
	fmt.Println(QuickSort(arr3))
}

func QuickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	splitdata := arr[0]          //第一个数据
	low := make([]int, 0, 0)     //比我小的数据
	hight := make([]int, 0, 0)   //比我大的数据
	mid := make([]int, 0, 0)     //与我一样大的数据
	mid = append(mid, splitdata) //加入一个
	for i := 1; i < len(arr); i++ {
		if arr[i] < splitdata {
			low = append(low, arr[i])
		} else if arr[i] > splitdata {
			hight = append(hight, arr[i])
		} else {
			mid = append(mid, arr[i])
		}
	}
	low, hight = QuickSort(low), QuickSort(hight)
	myarr := append(append(low, mid...), hight...)
	return myarr
}

//冒泡排序
func BubbleSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr
}
