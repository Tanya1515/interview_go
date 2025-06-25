package main

import (
	"testing"
	"unsafe"
)

// При конвертации строк происходит дополнительная аллокация
// памяти, так как строки неизменяемые в отличии от срезов,
// они должны шарить разные учатски памяти. Однако, можно
// производить конвертация строк без аллокаций.

// go test -bench=. -benchmem comparison_test.go

// Функция Convert преобразует слайс байт в строку
// без дополнительных аллокаций памяти.
func Convert(data []byte) string {
	if len(data) == 0 {
		return ""
	}

	// unsafe.SliceData - функция, которая возвращает указатель на массив
	// unsafe.String - функция, которая конструирует строку
	return unsafe.String(unsafe.SliceData(data), len(data))
}

var Result string

func BenchmarkConvertion(b *testing.B) {
	slice := []byte("Hello world!!!")
	for i := 0; i < b.N; i++ {
		Result = string(slice)
	}
}

func BenchmarkUnsafeConvertion(b *testing.B) {
	slice := []byte("Hello world!!!")
	for i := 0; i < b.N; i++ {
		Result = Convert(slice)
	}
}
