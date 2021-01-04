package main

import (
	"container/list"
	"fmt"
)

func main() {
	var emptyList1 = list.List{} // {{<nil> <nil> <nil> <nil>} 0}
	var emptyList2 list.List     // {{<nil> <nil> <nil> <nil>} 0}
	emptyList3 := list.New()     // &{{0xc00005c150 0xc00005c150 <nil> <nil>} 0}
	fmt.Println(emptyList1, emptyList2, emptyList3)

	l := list.New()
	e4 := l.PushBack(4)
	e1 := l.PushFront(1)
	l.InsertBefore(3, e4)
	l.InsertAfter(2, e1)
	fmt.Printf("%v\n", *l) // {{0xc00005c270 0xc00005c240 <nil> <nil>} 4}

	// 前序
	for i := l.Front(); i != nil; i = i.Next() {
		fmt.Printf("%-3v\n", i.Value) // 1  2  3  4
	}

	// 后序
	for i := l.Back(); i != nil; i = i.Prev() {
		fmt.Printf("%-3v\n", i.Value) // 4  3  2  1
	}

}
