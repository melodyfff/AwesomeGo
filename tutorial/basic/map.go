package main

import (
	"fmt"
	"sync"
)

// map 读线程安全，写则需要 sync.Map 保证线程安全

func main() {
	// var map_variable map[key_data_type] value_data_type = make(map[key_data_type] value_data_type , size)
	var m1 map[string]string = make(map[string]string)
	m1["1"] = "test1"
	m1["2"] = "test2"
	fmt.Println(m1) // map[1:test1 2:test2]

	m2 := map[string]string{"1": "test1", "2": "test2"}

	// range for
	for _, val := range m1 {
		fmt.Print(val) // test1test2
	}

	for key, _ := range m1 {
		fmt.Print(key) // 12
	}

	val, flag := m2["2"]
	fmt.Printf("\n%v - %v", flag, val) // true - test2

	delete(m2, "2")
	fmt.Println(m2) // map[1:test1]

	// sync.Map 无需初始化
	var syncMap sync.Map
	syncMap.Store("ok", "ok")
	syncMap.Store("ok2", "ok2")

	fmt.Printf("%v\n", &syncMap)
	fmt.Println(syncMap.Load("ok")) // ok true

}
