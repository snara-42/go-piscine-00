package main

import "fmt"

func subSlice(slice []int, begin int, length int, capacity int) (s []int) {
	if capacity < length {
		capacity = length
	}
	s = make([]int, length, capacity)
	if begin < len(slice) {
		copy(s, slice[begin:])
	}
	return s
}

func main() {
	var orig = []int{0, 1, 2, 3, 4, 5}
	var ret []int
	ret = subSlice(orig, 0, 3, 3)
	fmt.Printf("ret = %v, len = %d, cap = %d\n", ret, len(ret), cap(ret))
	ret = subSlice(orig, 2, 7, 10)
	fmt.Printf("ret = %v, len = %d, cap = %d\n", ret, len(ret), cap(ret))
	ret = subSlice(orig, 2, 7, 3)
	fmt.Printf("ret = %v, len = %d, cap = %d\n", ret, len(ret), cap(ret))
}
