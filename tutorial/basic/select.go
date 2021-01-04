package main

import (
	"fmt"
	"time"
)

func sender1(ch chan string) {
	for {
		time.Sleep(3 * time.Second)
		ch <- "ok from sender1"
	}
}

func sender2(ch chan string) {
	for {
		time.Sleep(1 * time.Second)
		ch <- "ok from sender2"
	}
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go sender1(ch1)
	go sender2(ch2)

	for {
		select {
		case r1 := <-ch1:
			fmt.Printf("r1 Accept -> %v\n", r1)
		case r2 := <-ch2:
			fmt.Printf("r2 Accept -> %v\n", r2)
		case <-time.After(4 * time.Second):
			fmt.Println("Time Out...")
		}
	}
}
