package main

import (
	"fmt"
)

// Здесь есть два пустых интерфейса,
// в которых хранится число 3.

func main() {
	var x interface{} = 3
	var y interface{} = 3
	fmt.Println(x == y) // true

	// var x interface{} = 3
	// var y interface{} = 3.0
	// fmt.Println(x == y) // false, поскольку разные типы объектов, которые стоят за пустыми интерфейсами
}
