package main

import (
	"fmt"
	"time"
)

func cost() {
	start := time.Now()
	time.Sleep(time.Duration(1) * time.Second)
	end := time.Since(start)
	fmt.Println("run ", end)
}

func main() {
	cost()
}
