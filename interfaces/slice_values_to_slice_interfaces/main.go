package main

import "fmt"

func main() {

	// Срез пустых интерфейсов
	var interfaces []interface{}

	// Срез целочисленных значений
	values := []int{100, 200, 300, 400, 500}

	// В языке golang операцию присваивания среза пустых интерфейсов
	// срезу целочисленных элементов присваивать нельзя.

	interfaces = values

	/*
		Однако, присвоение можно выполнять поэлементно.

		interfaces = make([]interface{}, len(values))
		for idx := 0; idx < len(values); idx++ {
			interfaces[idx] = values[idx]
		}
	*/

	fmt.Println(interfaces...)
}
