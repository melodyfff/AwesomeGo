package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

const (
	UdpAddress = "127.0.0.1:8081"
	UdpNetwork = "udp"
)

func handleErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	// resolve address
	addr, err := net.ResolveUDPAddr(UdpNetwork, UdpAddress)

	handleErr(err)

	// listen
	conn, err := net.ListenUDP(UdpNetwork, addr)

	handleErr(err)

	defer conn.Close()

	fmt.Println("[UDP] Listen to addr - ", addr)

	for {
		// cache
		data := make([]byte, 1024)

		// block
		_, rAddr, err := conn.ReadFromUDP(data)

		if err != nil {
			fmt.Println("Read From Udp err = ", err)
			continue
		}

		strData := string(data)
		fmt.Println("Received: ", strData)

		upper := strings.ToUpper(strData)
		_, err = conn.WriteToUDP([]byte(upper), rAddr)
		if err != nil {
			fmt.Println("Write Upd err = ", err)
			continue
		}

		fmt.Println("Send: ", upper)
	}
}
