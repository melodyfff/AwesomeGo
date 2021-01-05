package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("udp", "127.0.0.1:8081")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer conn.Close()

	// input from stdin
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		line := input.Text()
		if _, err := conn.Write([]byte(line)); err != nil {
			fmt.Println("Write err = ", err)
			return
		}
		fmt.Println("Write: ", line)

		// read from server
		msg := make([]byte, 1024)
		if _, err := conn.Read(msg); err != nil {
			fmt.Println("Read err = ", err)
			return
		}

		fmt.Println("Response: ", string(msg))
	}

}
