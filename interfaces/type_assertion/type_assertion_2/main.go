package main

import "fmt"

// как работает type assertion:
// x.(T) проверяет, что x != nil и:
// 1) если T не интерфейс, то проверяет,
// что динамический тип x - это T
// 2) если T - интерфейс, то проверяет,
// что динамический тип х его реализует.

// По этой причине код ниже работает.

type fooer interface{ foo() }
type barer interface{ bar() }
type foobarer interface {
	foo()
	bar()
}

type thing struct{}

func (t *thing) foo() {}
func (t *thing) bar() {}

func main() {
	var i foobarer = &thing{}
	_, ok := i.(fooer)
	fmt.Println("result:", ok) // true
}
