package main

import "testing"

// Есть несколько способов, которые позволяют модифицировать значения 
// внутри массива. 

// go test -bench=. comparison_test.go

// Есть некоторая структура. 
type account struct {
	balance int
}

func BenchmarkWithPointers(b *testing.B) {
	// массив accounts может хранить не просто значения
	// а указатели на стрктуры. 
	accounts := [...]*account{
		{balance: 100},
		{balance: 200},
		{balance: 300},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, a := range accounts {
			// здесь будет происходить обращение по указателю
			a.balance += 1
		}
	}
}

func BenchmarkWithIndices(b *testing.B) {
	accounts := [...]account{
		{balance: 100},
		{balance: 200},
		{balance: 300},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for i := range accounts {
			// здесь модификация происходит по индексу 
			// то есть есть некоторое начало массива - accounts,
			// известно, какое количество данных в массиве,
			// а также их тип. Таким образом, при обращении к элементу по индексу
			// происходит обращение в память: accounts + i*(размер типа элемента)
			// Причем этот способ будет работать быстрее, поскольку гораздо проще
			// итеративно ходить по памяти, чем прыгать по указателям в памяти. 
			accounts[i].balance += 1
		}
	}
}
