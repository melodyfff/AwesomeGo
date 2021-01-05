package main

import (
	"fmt"
	"time"
)

// 接收一个只写的channel
func producer(send chan<- string) {
	defer close(send)
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second)
		send <- "ok"
		fmt.Println("产生数据")
	}
}

// 接收一个只读的channel
func consumer(receive <-chan string) {
	for msg := range receive {
		time.Sleep(3 * time.Second)
		fmt.Printf("接收数据 msg = '%v'\n", msg)
	}
}

func handleChannelTimeOut() {
	// 无缓冲channel
	ch := make(chan string)
	quit := make(chan bool)

	go func() {
		for {
			select {
			case s := <-ch:
				fmt.Printf("Accept -> %v \n", s)
			case <-time.After(3 * time.Second):
				fmt.Println("Channel Time Out...")
				quit <- true
			}
		}
	}()

	for i := 0; i < 10; i++ {
		ch <- "ok"
		time.Sleep(time.Second)
	}

	// wait
	<-quit

	fmt.Println("Over.")

}

func main() {
	var word chan string     // <nil>
	word = make(chan string) // 无缓冲
	go func() {
		time.Sleep(1 * time.Second)
		word <- "Hello World!"
	}()
	fmt.Println("waiting...")
	msg, isAccept := <-word
	fmt.Printf("msg = '%v' ,accept data ? %v \n ", msg, isAccept)

	// 带缓冲的channel
	ch := make(chan interface{}, 10)
	ch <- "Hello"
	ch <- "World"
	ch <- "!"
	close(ch)
	for {
		// isAccept表示是否接收到数据
		if msg, isAccept := <-ch; isAccept == false {
			fmt.Printf("msg = '%v' ,accept data ? %v \n ", msg, isAccept)
			break
		} else {
			fmt.Printf("msg = '%v' ,accept data ? %v \n ", msg, isAccept)
		}
	}

	// 模拟通道发射信息， range阻塞等待接收数据
	ch = make(chan interface{})
	go func() {
		defer close(ch) // finally close ch.
		for i := 0; i < 5; i++ {
			ch <- "ok."
			time.Sleep(time.Second)
		}
	}()
	for data := range ch {
		fmt.Printf("data = '%v'\n", data)
	}

	// 生产者消费者 - 无缓冲 , 无缓冲的情况下只有当goroutine都准备好了才继续，否则阻塞。
	// 这里producer sleep 1s 而 consumer sleep 3s ，使用无缓冲则会造成等待consumer准备好才继续
	channel := make(chan string)
	go producer(channel)
	consumer(channel)

	handleChannelTimeOut()
}
