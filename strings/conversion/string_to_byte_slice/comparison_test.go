package main

import (
	"testing"
	"unsafe"
)

// Здесь представлена обратная операция по конвертации
// строки в срез. Здесь аналогично преобразования происходят
// при помощи пакета unsafe.

// go test -bench=. -benchmem comparison_test.go

func Convert(str string) []byte {
	if len(str) == 0 {
		return nil
	}

	return unsafe.Slice(unsafe.StringData(str), len(str))
}

var Result []byte

func BenchmarkConvertion(b *testing.B) {
	str := "Hello world!!!"
	for i := 0; i < b.N; i++ {
		Result = []byte(str)
	}
}

func BenchmarkUnsafeConvertion(b *testing.B) {
	str := "Hello world!!!"
	for i := 0; i < b.N; i++ {
		Result = Convert(str)
	}
}
