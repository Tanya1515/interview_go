package main

import (
	"fmt"
	"unsafe"
)

type data struct{}

func main() {
	x := &data{} // x = &y - эквивалентно
	y := data{}

	fmt.Println(unsafe.Sizeof(x)) // 8 - размер машинного слова
	fmt.Println(unsafe.Sizeof(y)) // 0
}
