package main

/* 定义接口 */
// type interface_name interface {
//		method_name [return_type]
// }

/* 定义结构体 */
// type struct_name struct {
//	  	/* variables */
// }

/* 实现接口方法 */
// func (struct_name_variable struct_name) method_name1() [return_type] {
//		/* 方法实现 */
// }

import "fmt"

type Coffee interface {
	make()
}

type Coffee1 struct {
}

func (c Coffee1) make() {
	fmt.Println("make coffee1.")
}

func main() {
	coffee := new(Coffee1)
	coffee.make()
}
