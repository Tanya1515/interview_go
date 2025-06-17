package main

import "fmt"

type Builder struct {
	buffer []byte
}

// Это примерная реализация strings.Builder-а

// Конструктор
func NewBuilder() Builder {
	return Builder{}
}

func (b *Builder) Grow(capacity int) {
	if capacity < 0 {
		return
	}

	// Если требуемый размер меньше текущей длины среза
	// то отрезаем кусок необходимой длины
	if capacity < len(b.buffer) {
		b.buffer = b.buffer[:capacity]
		return
	}

	// Если размерность больше текущей длины -
	// создаем отдельный буфер с требуемой capacity
	// Далее перекопируем старые значения в новый буфер
	// и сохраним выделенный буфер в структуру Builder
	buffer := make([]byte, len(b.buffer), capacity)
	copy(buffer, b.buffer)
	b.buffer = buffer
}

func (b *Builder) Write(symbol byte) {
	b.buffer = append(b.buffer, symbol)
}

// Метод At - метод, который возвращает указатель
// на byte по индексу.
func (b *Builder) At(index int) *byte {
	if index < 0 || index >= len(b.buffer) {
		return nil
	}

	return &b.buffer[index]
}

func (b *Builder) String() string {
	return string(b.buffer)
}

func main() {
	builder := NewBuilder()
	builder.Grow(3)

	builder.Write('a')
	builder.Write('b')
	builder.Write('c')

	fmt.Println(builder.String())
}
