package main

import "fmt"

type fooer interface{ foo() }

type thing struct{}

func (t *thing) foo() {}
func (t *thing) bar() {}

func main() {
	var i fooer = &thing{}

	// dynamically checking for certain methods
	_, ok := i.(interface{ bar() })

	// type assertion проверяет динамический тип -
	// это thing, а интерфейс fooer - это статический тип.
	// При этом у thing есть метод bar, поэтому он соответствует
	// интерфейсу, к которому хотим привести i.
	fmt.Println("result:", ok) // true
}
