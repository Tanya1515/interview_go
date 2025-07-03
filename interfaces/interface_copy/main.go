package main

import (
	"fmt"
	"unsafe"
)

type eface struct {
	typ unsafe.Pointer
	val unsafe.Pointer
}

func main() {
	// Здесь представлено некоторое целочисленное значение,
	// которое записывается в пустой интерфейс.
	var value int = 100
	var i interface{} = value // при присвоении некоторого значения интерфейсу - оно копируется
	fmt.Println(i)            // 100

	// Далее меняем значение, которое находится в пустом интерфейсе
	value = 200
	fmt.Println(i) // 100 - значение не изменилось

	/*
		При запуске данного кода будет видно, что obj и value,
		которое записывается в obj указывает на разные участки
		памяти.

		Здесь приводится указатель к obj.
		obj := (*eface)(unsafe.Pointer(&i))
		println("&value:", &value)
		println("obj.val:", obj.val)
	*/
}
