package main

import "fmt"

func main() {
	// var identifier []type = make([]T, length, capacity)

	var slice1 = make([]int, 2, 5) // len=2 cap=5 slice=[0 0]

	arr := []int{1, 2, 3}
	slice2 := arr[1:] // len=2 cap=2 slice=[2 3]

	slice3 := arr[:]
	slice3 = append(slice3, 4, 5) // len=5 cap=6 slice=[1 2 3 4 5]

	fmt.Printf("len=%d cap=%d slice=%v\n", len(slice1), cap(slice1), slice1)
	fmt.Printf("len=%d cap=%d slice=%v\n", len(slice2), cap(slice2), slice2)
	fmt.Printf("len=%d cap=%d slice=%v\n", len(slice3), cap(slice3), slice3)
}
