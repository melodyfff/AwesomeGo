package main

import "fmt"

// 多个变量
func func1(n1, n2 string) int {
	// n1,n2 is string
	return 1
}

// 多个返回
func func2(n1 int32, n2 string) (string, int) {
	return "ok", 1
}

// 不定参数
func func3(args ...string) {
	for _, val := range args {
		fmt.Println(val)
	}
}

// 修改传入参数的真实值
func func4(val *int) {
	// **int 为真实的值
	*val += 1
}

// 匿名函数
func anonymous() {
	// 匿名函数
	func() {
		fmt.Println("ok")
	}()

	// 带传入和传出的匿名函数
	func(n1, n2 int) {
		fmt.Println(n1, n2)
	}(10, 20)
}

var anonymous1 = func() {
	fmt.Println("ok")
}

var anonymous2 = func(arg1 int) string {
	return "ok"
}
var anonymous3 = func(args ...int) (int, int) {
	return 1, 2
}

// 匿名函数作为参数传入
func anonymous4(arg1 int, callback func(int) string) {
	callback(arg1)
}

// 闭包函数 - 自增
func add() func() int {
	var n int = 0
	return func() int {
		n += 1
		return n
	}
}

func main() {
	// func funcName(paramlist paramType, ...) (returnType, ...){ }
	fmt.Println()
	// 匿名函数使用实例
	anonymous() // ok 10 20

	anonymous1() // ok

	anonymous4(1, anonymous2) // ok

	c := func(num1, num2 int) int {
		return num1 + num2
	}
	fmt.Println(c(1, 2)) // 3

	// 闭包自增
	a := add()
	fmt.Println(a()) // 1
	fmt.Println(a()) // 2
	fmt.Println(a()) // 3
}
