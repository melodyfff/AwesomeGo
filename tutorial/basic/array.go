package main

import "fmt"

func main() {
	// var variable_name [SIZE] variable_type
	var array = [5]int32{1, 2, 3, 4}   // [1 2 3 4 0]
	var array2 = []float32{1, 2, 3, 4} // [1 2 3 4]
	array3 := []int64{1}               // [1]
	var empty []string                 // []

	fmt.Println(array, array2, array3, empty)
	fmt.Println(len(array), len(array2), len(array3), len(empty))

	// fori
	for i := 0; i < len(array2); i++ {
		fmt.Print(array2[i]) // 1234
	}

	// range for
	for _, val := range array2 {
		fmt.Print(val) // 1234
	}

	for index := range array2 {
		fmt.Print(index) // 0123
	}
}
