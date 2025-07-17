package main

import (
	"sync"
)

/*

	Проблемы кода:

	1) Нельзя копировать примитивы синхронизации. А здесь в структуре Stack передается копия mutex.
		Соответственно, можно либо передавать указатель на mutex, либо в качестве reciever
		передавать указатель на структуру Stack.

	2) В коде присутствует data race. Функция append добавляет элемент в слайс,
	а затем инкрементирует его размер. Однако, в методе Pop/Top происходит чтение длины слайса.
	При этом проверку длины можно выполнять под mutex-ом.

	3) Может быть проблема с некорректным чтением и доставанием элемента из стека.

	[1,2]
	G1 Top() -> 2
					G2 Top() -> 2
	G1 Pop() -> 2 [1]
					G2 Pop() -> 1 []

	Поэтому чаще всего делают один метод, который одновременно возвращает значение и удаляет элемент из стека.
*/

// Есть структура Stack, в которой есть mutex,
// где синхронизируется доступ к слайсу.
type Stack struct {
	mutex sync.Mutex
	data  []string
}

func NewStack() Stack {
	return Stack{}
}

// Метод Push, который добавляет в stack элемент.
func (b Stack) Push(value string) {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	b.data = append(b.data, value)
}

// Метод Pop, который снимает элемент из стека.
func (b Stack) Pop() {
	if len(b.data) == 0 {
		panic("pop: stack is empty")
	}

	b.mutex.Lock()
	defer b.mutex.Unlock()

	b.data = b.data[:len(b.data)-1]
}

// Метод Top, который возвращает элемент из стека.
func (b Stack) Top() string {
	if len(b.data) == 0 {
		panic("top: stack is empty")
	}

	b.mutex.Lock()
	defer b.mutex.Unlock()

	return b.data[len(b.data)-1]
}

var stack Stack

// Есть функция producer, которая push-ит 1000 элементов в стек.
func producer() {
	for i := 0; i < 1000; i++ {
		stack.Push("message")
	}
}

// Есть функция consumer которая сначала достает элемент
// с верхушки стека, а затем удаляет этот элемент из стека.
// При этом достают только 10 элементов из стека.
func consumer() {
	for i := 0; i < 10; i++ {
		_ = stack.Top()
		stack.Pop()
	}
}

func main() {
	producer()

	wg := sync.WaitGroup{}
	wg.Add(100)

	// Здесь consumer вызывается 100 раз, соответственно,
	// по завершении работы цикла достанем все 1000 сообщений.
	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			consumer()
		}()
	}

	// а затем дожидаемя выполнения всех горутин
	wg.Wait()
}
