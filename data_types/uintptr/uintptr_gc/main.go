package main

import (
	"runtime"
	"unsafe"
)

// go run main.go
// go run -gcflags=-d=checkptr main.go

func main() {
	x := new(int)
	y := new(int)
	z := new(int)

	ptrX := unsafe.Pointer(x)
	ptrY := unsafe.Pointer(y)
	addressZ := uintptr(unsafe.Pointer(z)) // адрес переменной в виде целочисленного значения

	// arithmetic operation
	_ = addressZ + 2
	_ = addressZ - 2

	// после вызова garbage collector-а  erfpfntkm я  может быть очищен,
	// поскольку uintptr - это просто число,
	// а сборщик мусора трассирующий, то есть он смотрит на указатели
	// а uintptr не хранит указатель на некоторый объект, поэтому z может быть очищен
	runtime.GC()

	*(*int)(ptrX) = 100
	*(*int)(ptrY) = 200
	*(*int)(unsafe.Pointer(addressZ)) = 300 // dangerous
}
