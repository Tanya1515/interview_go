package main

import (
	"strings"
	"testing"
	"unsafe"
)

// Смысл в том, что когда создается slice,
// то в нем указывается и его длина, и capacity
// Стоит учитывать, что все элементы в
// созданном слайсе будут проинициализированы
// zerovalue-значениями. Например, для int-ового слайса
// - это нули. Однако, возникают случаи,
// когда это избыточно.

// go test -bench=. creation_test.go

func makeDirty(size int) []byte {

	// в следующих двух строчках кода аллоцируется
	// буфер размера size

	var sb strings.Builder
	sb.Grow(size)

	// При помощи пакета unsafe получаем указатель
	// на начало этого буфера. А затем возвращаем слайс.
	pointer := unsafe.StringData(sb.String())
	return unsafe.Slice(pointer, size)
}

var Result []byte

func BenchmarkMake(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Result = make([]byte, 0, 10<<20)
	}
}

func BenchmarkMakeDirty(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Result = makeDirty(10 << 20)
	}
}
