package main

import (
	"fmt"
	"unsafe"
)

// Строки в golang неизменяемые, однако, их можно модифицировать
// при помощи пакета unsafe.

func main() {
	// Срез data и строка strData шарят
	// один и тот же участок памяти
	data := []byte("Hello world")
	strData := unsafe.String(unsafe.SliceData(data), len(data))

	fmt.Println(strData) // Hello world
	data[0] = 'W'
	fmt.Println(strData) // Wello world
}
