package main

import "fmt"

func main() {
	// var var_name *var-type
	var num int = 10
	var p1 *int = &num // 0xc00001a0c0
	p2 := &num         // 0xc00001a0c0

	var null *string // <nil>

	fmt.Println(p1, p2, null)
	fmt.Println(*p1, *p2)
	fmt.Printf("null -> %x", null) // 0
}
