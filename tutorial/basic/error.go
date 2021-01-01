package main

import (
	"errors"
	"fmt"
)

var (
	myError    = errors.New("this is a error")
	panicError = errors.New("this is a panic error")
)

func errorMaker(flag bool) (int, error) {
	if flag {
		return 0, myError
	}
	return 0, nil
}

func errorMaker2(flag bool) (int, error) {
	if flag {
		panic(panicError) // 终止程序运行
	} else {
		return 0, nil
	}
}

func main() {
	defer func() {
		// recover()处理异常，恢复程序运行
		if err := recover(); err != nil {
			fmt.Printf("recover(): %v", err)
		}
	}()

	fmt.Println(errorMaker(false)) // 0 <nil>
	fmt.Println(errorMaker(true))  // 0 this is a error!

	fmt.Println(errorMaker2(false)) // 0 <nil>
	fmt.Println(errorMaker2(true))  // panic: this is a panic error! 有recover()则在recover中处理
}
