package main

import (
	"strings"
	"testing"
)

// Для ускорения конкатенации строк используется специальная конструкция -
// strings.Builder. Для сложения строк есть специальный метод WriteString
// В конце значение внутри builder-а приводится к строке. В случае если
// понятно, какой итоговый размер будет у строки, можно использовать метод
// Grow, чтобы заранее зарезервировать необходимое количество памяти.
// Самым эффективным способом для конкатенации строк будет пример с
// ранее выделенным количеством памяти под строку.

// go test -bench=. speed_test.go

func BenchmarkSimpleConcatenation(b *testing.B) {
	str := "test"
	for i := 0; i < b.N; i++ {
		str += "test"
	}

	_ = str
}

func BenchmarkConcatenationWithStringBuilder(b *testing.B) {
	builder := strings.Builder{}
	builder.WriteString("test")
	for i := 0; i < b.N; i++ {
		builder.WriteString("test")
	}

	_ = builder.String()
}

func BenchmarkConcatenationWithStringBuilderOptimized(b *testing.B) {
	builder := strings.Builder{}
	builder.Grow(4 + b.N*4)
	builder.WriteString("test")
	for i := 0; i < b.N; i++ {
		builder.WriteString("test")
	}

	_ = builder.String()
}
