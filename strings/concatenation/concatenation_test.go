package main

import (
	"fmt"
	"strings"
	"testing"
)

// В этом варианте самым быстрым способом является конкатенация строк
// при помощи оператора +. Sprintf работает крайне долго, поскольку
// под капотом используется рефлексия, Join работает медленнее примерно
// в два раза. Конкатенация с использованием оператора + оптимизирована
// с точки зрения аллокации памяти. В этом случае вычисляется суммарный
// размер при конкатенации, будет сделана только одна аллокация, куда
// будут скопированы все строки.

// go test -bench=. concatenation_test.go

func BenchmarkConcatenationWithSprintf(b *testing.B) {
	s0 := "str1"
	s1 := "str2"
	s2 := "str3"
	s3 := "str4"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("%s%s%s%s", s0, s1, s2, s3)
	}
}

func BenchmarkConcatenationWithOperatorPlus(b *testing.B) {
	s0 := "str1"
	s1 := "str2"
	s2 := "str3"
	s3 := "str4"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = s0 + s1 + s2 + s3
	}
}

func BenchmarkConcatenationWithJoin(b *testing.B) {
	s0 := "str1"
	s1 := "str2"
	s2 := "str3"
	s3 := "str4"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = strings.Join([]string{s0, s1, s2, s3}, "")
	}
}
