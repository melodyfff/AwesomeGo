package main

import "fmt"

type (
	Commodity struct {
		// 变量名小写为私有访问 private
		name  string
		price int
	}

	Order struct {
		// public
		Id int
	}

	Aggregation struct {
		Commodity Commodity
		Order     Order
	}
)

// method - 接收器 (c Commodity)
func (c Commodity) print() {
	fmt.Printf("%#v\n", c)
}

// 指针接收器
func (o *Order) print() {
	fmt.Printf("%#v\n", o)
}

func main() {
	// 声明
	var emptyCommodity Commodity
	emptyCommodity.print() // main.Commodity{name:"", price:0}

	// 声明并赋值
	initCommodity := Commodity{"test", 10}
	fmt.Printf("%+v\n", initCommodity) // {name:test price:10}

	var order = new(Order) // 分配内存
	order.print()          // &main.Order{Id:0}
	var order2 = &Order{1}
	fmt.Printf("%v\n", order2) // &{1}

	// 匿名结构图
	anonymous := struct {
		name string
	}{
		name: "ok",
	}
	fmt.Printf("%+v\n", anonymous) // {name:ok}

	// 聚合
	aggregation := Aggregation{
		Commodity: Commodity{name: "test2", price: 20},
		Order:     Order{1},
	}
	fmt.Printf("%+v\n", aggregation) // {Commodity:{name:test2 price:20} Order:{Id:1}}
}
