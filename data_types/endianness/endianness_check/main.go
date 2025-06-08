package main

import (
	"fmt"
	"unsafe"
)

// Код, который позволяет проверить, в каком формате записано
// число: Little/Big Endian.

func IsLittleEndian() bool {
	// выбираем число, которое состоит больше, чем из одного байта. 
	var number int16 = 0x0001
	// берем адрес на начало этого числа, приводим к однобайтовому значению
	// другими словами теперь pointer ссылается на первую часть числа. 
	pointer := (*int8)(unsafe.Pointer(&number))
	// разыменовываем указатель и сравниваем его с 1, поскольку если число
	// хранится в формате Little endian (0100), то в первом байте будет 01. 
	return *pointer == 1
}

func IsBigEndian() bool {
	return !IsLittleEndian()
}

func main() {
	if IsLittleEndian() {
		fmt.Println("Little endian")
	} else {
		fmt.Println("Big endian")
	}
}
