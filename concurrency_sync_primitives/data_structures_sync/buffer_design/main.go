package main

import "sync"

// Структура Bufer
type Buffer struct {
	mtx  sync.Mutex
	data []int
}

// Конструктор для создания Bufer-а
func NewBuffer() *Buffer {
	return &Buffer{}
}

// Добавляется элемент в Bufer
func (b *Buffer) Add(value int) {
	b.mtx.Lock()
	defer b.mtx.Unlock()

	b.data = append(b.data, value)
}

// Возвращается срез данных из буфера.
// Вот здесь проблема, поскольку срез
// представляет собой структуру данных,
// которая включает в себя длину, размерность
// и указатель на начало среза. Причем
// мы возвращаем копию этой структуры,
// которая может модифицироваться дальше по коду
// и изменяться, при этом в несинхронизированном режиме.
// То есть в другом месте может происходить операция
// записи/чтения.

func (b *Buffer) Data() []int {
	b.mtx.Lock()
	defer b.mtx.Unlock()

	return b.data
}

/*
	Существуют два способа решения этой проблемы:

	1) Можно явно копировать слайс, используя DeepCopy

	2) Можно передавать внутрь функцию, которая будет
	обращаться к каждому элементу слайса. То есть здесь
	был реализован некоторый callback, который обходит срез
	с синхронизацией.

	func (b *Buffer) ForEach(action func(int)) {

		if action == nil {
			return
		}
		b.mtx.Lock()
		defer b.mtx.Unlock()

		for _, value := b.data {
			action(value)
		}
	}

*/
