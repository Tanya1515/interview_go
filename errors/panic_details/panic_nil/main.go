package main

import (
	"fmt"
	"runtime"
)

/*
Паника принимает пустой интерфейс. То есть туда можно пробросить
все, что угодно, - 1, строку и так далее.
*/
func panicNil() {
	defer func() {
		fmt.Println("recovered:", recover())
	}()

	panic(nil)
}

func panicNilError() {
	defer func() {
		fmt.Println("recovered:", recover())
	}()

	panic(new(runtime.PanicNilError))
}

// Вывод:
// recovered: panic called with nil argument
// recovered: panic called with nil argument

// и в 1ом, и во 2ом случаях паника вызвана с nil-ым аргументом

func main() {
	panicNil()
	panicNilError()
}
