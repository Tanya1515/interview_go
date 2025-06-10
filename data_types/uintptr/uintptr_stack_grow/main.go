package main

import (
	"fmt"
	"unsafe"
)

//go:noinline
func allocation(index int) byte {
	// здесь выделяется массив на 10 Mb
	var data [10 << 20]byte
	return data[index]
}

func main() {
	var array [10]int

	// здесь берется адрес массива в целочисленном представлении
	address1 := (uintptr)(unsafe.Pointer(&array))
	fmt.Println("#1 array address:", address1)

	// после вызова этой функции сильно вырастет размер стека, поэтому
	// причем размер стека растет также как и слайс, то есть память переезжает
	// поэтому массив array скорее всего переедет в другое место
	allocation(100)

	address2 := (uintptr)(unsafe.Pointer(&array))
	// переменные address1 и address2 будут иметь разные значения
	// поскольку при росте стека указатели будут обновляться, а uintptr - нет
	fmt.Println("#2 array address:", address2)
	fmt.Println("#1 array address:", address1)
}
