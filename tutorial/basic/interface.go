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

import (
	"fmt"
	"reflect"
)

type (
	Coffee interface {
		make()
	}

	Cooker interface {
		Coffee
		work()
	}

	CoffeeName struct {
		name string
	}
)

// implement
func (c CoffeeName) make() {
	fmt.Printf("%#v make coffee1.\n", c)
}

func (c *CoffeeName) work() {
	fmt.Printf("%#v work coffee1.\n", c)
}

func main() {
	coffee := new(CoffeeName)
	coffee.make()
	coffee.work()
	fmt.Println(reflect.TypeOf(coffee)) // *main.CoffeeName

	var coffee1 interface{} = &CoffeeName{"ok"}
	fmt.Println(reflect.TypeOf(coffee1)) // *main.CoffeeName
	// 同时是 Coffee类型也是Cooker类型
	fmt.Println(reflect.TypeOf(coffee1.(Cooker)))
	fmt.Printf("%#+v\n", coffee1.(Coffee)) // &{name:ok}
	fmt.Printf("%#+v\n", coffee1.(Cooker)) // &{name:ok}

	// 空接口，可放任何值 . 空接口是任何类型的父接口
	var x interface{}
	x = 100
	x = "ok"
	x = true
	fmt.Println(x)                 // true
	fmt.Println(reflect.TypeOf(x)) // boolean

	// value,ok := x.(T) 语言类型断言 ， 如果使用bool变量接收，失败时会直接 panic() 	panic: interface conversion: interface {} is bool, not int
	if v, ok := x.(bool); ok {
		fmt.Println(v) // true
	}

}
