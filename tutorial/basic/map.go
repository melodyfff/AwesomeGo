package main

import "fmt"

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
}
