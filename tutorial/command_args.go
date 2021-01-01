package main

// 命令行传参解析
// go run command_args.go -u ok -p 1
// go run command_args.go --help

import (
	"flag"
	"fmt"
	"os"
)

func arg1() {
	var user string
	var password string
	flag.StringVar(&user, "u", "root", "用户名,默认root")
	flag.StringVar(&password, "p", "123", "密码,默认123")
	// 解析命令行参数
	flag.Parse()
	// go run command_args.go -u ok -p 1
	fmt.Printf("\nuser: %v \t password:%v\n", user, password)
}

func arg2() {
	var user *string
	var password *string
	user = flag.String("u", "root", "用户名,默认root")
	password = flag.String("p", "123", "密码,默认123")
	// 解析命令行参数
	flag.Parse()
	// go run command_args.go -u ok -p 1
	fmt.Printf("\nuser: %v \t password:%v\n", *user, *password)
}

func main() {
	args := os.Args
	fmt.Printf("args: %v", args)
	//arg1()
	arg2()
}
