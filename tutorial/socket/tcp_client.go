package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	fmt.Println("Client start connect to 0.0.0.0:8080")
	conn, err := net.Dial("tcp", "0.0.0.0:8080")
	if err != nil {
		fmt.Println("connect failed ,err :", err)
		return
	}

	defer conn.Close()

	input := bufio.NewReader(os.Stdin)
	for {
		in, err := input.ReadString('\n')
		if err != nil {
			fmt.Println("Read from console failed, err: ", err)
			break
		}

		data := strings.TrimSpace(in)
		if data == "Q" || data == "q" {
			break
		}

		// send data
		if _, err := conn.Write([]byte(data)); err != nil {
			fmt.Println("Write failed, err: ", err)
			break
		}

		// rec
		recData := make([]byte, 1024)
		if _, err := conn.Read(recData); err != nil {
			fmt.Println("Read failed, err: ", err)
			break
		}
		fmt.Println("Receive data: ", string(recData))
	}

}
