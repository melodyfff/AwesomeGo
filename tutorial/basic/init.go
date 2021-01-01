package main

// 全局变量的定义 -> init()函数 -> main()

import "fmt"

func initArgs() string {
	fmt.Println("initArgs run.")
	return "ok111"
}

var arg = initArgs()

func init() {
	fmt.Println("init run.")
}

func main() {
	fmt.Println("ok")
	// initArgs run.
	// init run.
	// ok
}
