package main

import (
	"fmt"
	"net"
)

func process(c net.Conn) {
	defer c.Close()

	for {
		// cache
		var buf [1024]byte

		// Read
		n, err := c.Read(buf[:])
		if err != nil {
			fmt.Println("read from connect failed, err: ", err)
			break
		}
		fmt.Println("receive from client, data: ", string(buf[:n]))

		// Send
		if _, err := c.Write([]byte("Send From Server.")); err != nil {
			fmt.Println("write to client failed, err: ", err)
			break
		}
	}
}

func main() {
	fmt.Println("Server start listen to 0.0.0.0:8080")
	listener, err := net.Listen("tcp", "0.0.0.0:8080")

	if err != nil {
		fmt.Println("Listen fail ,err = ", err)
		return
	}

	for {
		// block
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("accept fail, err: ", err)
			continue
		}

		// 为每个请求新建携程处理
		go process(conn)
	}
}
