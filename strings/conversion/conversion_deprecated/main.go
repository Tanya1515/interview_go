package main

import (
	"fmt"
	"unsafe"
)

/*
type SliceHeader struct {
    Data unsafe.Pointer
    Len  int
    Cap  int
}

type StringHeader struct {
    Data unsafe.Pointer
    Len  int
}
*/

// Неправильный вариант конвертации, поскольку получается так
// что преобразуются первые два поля. Однако, если разработчики
// golang поменяют поля местами, то код будет стрелять по ногам
// (смотри на поля в структурах).

func ToStringDeprecated(data []byte) string {
	return *(*string)(unsafe.Pointer(&data))
}

func main() {
	data := []byte{'h', 'e', 'l', 'l', 'o'}
	str := ToStringDeprecated(data)
	fmt.Println(str)
}
