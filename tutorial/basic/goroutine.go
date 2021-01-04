package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	wg sync.WaitGroup
)

func main() {
	wg.Add(2) // 设置需要等待的携程数

	go func() { // go 标记goroutine
		defer wg.Done() // 携程结束

		for i := 0; i < 3; i++ {
			fmt.Println("1 RUN...")
			time.Sleep(time.Second)
		}
	}()

	go func() { // go 标记goroutine
		defer wg.Done() // 携程结束

		for i := 0; i < 3; i++ {
			fmt.Println("2 RUN...")
			time.Sleep(time.Second)
		}
	}()

	wg.Wait() // 等待所有携程结束

	fmt.Println("Main Down")
}
