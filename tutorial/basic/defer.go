package main

import "fmt"

func ok() {
	fmt.Println("ok1 run..")
	// defer 执行顺序先进后出
	defer fmt.Println("defer1 run.")
	defer fmt.Println("defer2 run.")
	fmt.Println("ok2 run..")
}

func main() {
	ok()
}
