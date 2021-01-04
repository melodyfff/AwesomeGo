package main

import (
	"fmt"
	"sync"
)

var (
	count int
	mutex sync.Mutex // 互斥锁
	group sync.WaitGroup
)

func increaseNoLock() {
	defer group.Done()
	count++
}

func increaseWithLock() {
	mutex.Lock() // lock

	defer mutex.Unlock()
	defer group.Done()

	count++
}

func main() {
	group.Add(100)
	for i := 0; i < 100; i++ {
		go increaseNoLock()
	}
	group.Wait()
	fmt.Println(count) // 无锁，此时count不一定为100

	count = 0 // 重置

	group.Add(100)
	for i := 0; i < 100; i++ {
		go increaseWithLock()
	}
	group.Wait()
	fmt.Println(count) // 互斥锁，结果必为 100
}
