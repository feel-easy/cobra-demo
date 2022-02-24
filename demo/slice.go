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
	fmt.Println(len(s))
}
