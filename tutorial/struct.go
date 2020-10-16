package main

import "fmt"

type Coffe struct {
	name  string
	price int
}

func main() {
	// 声明
	var emptyCoffee Coffe
	fmt.Printf("%#v\n", emptyCoffee)

	// 声明并赋值
	initCoffee := Coffe{"test", 10}
	fmt.Printf("%+v\n", initCoffee)
}
