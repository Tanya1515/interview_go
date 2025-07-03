package main

import "fmt"

type fooer interface{ foo() }
type barer interface{ bar() }
type foobarer interface {
	foo()
	bar()
}

type thing struct{}

func (t *thing) foo() {}
func (t *thing) bar() {}

var i foobarer = &thing{}

func main() {
	// Здесь type assetion успешно пройдет, как для
	// fooer, barer, foobarer, поскольку у типа thing
	// есть все необходимые методы.

	// При этом будет выбрана ветка, которая
	// располагается первой, даже если все остальные
	// ветки в true.

	switch v := i.(type) {
	case fooer:
		fmt.Println("fooer:", v)
	case barer:
		fmt.Println("barer:", v)
	case foobarer:
		fmt.Println("foobarer:", v)
	default:
		panic("none of them")
	}
}
