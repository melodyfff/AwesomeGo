package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// 闭包判断是否有该后缀名，没有则添加
func markSuffix(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func main() {
	m := markSuffix(".png")
	fmt.Println("FileName=", m("test"))     // FileName= test.png
	fmt.Println("FileName=", m("test.png")) // FileName= test.png

	// 读取当前路径下的文件
	dir, _ := ioutil.ReadDir("./")
	for _, i := range dir {
		fmt.Println(m(i.Name()))
	}
}
