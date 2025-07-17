package main

import (
	"runtime"
	"sync"
	"sync/atomic"
	"testing"
)

// Проблема в том, что когда счтечики были разнесены по элементам массива -
// предполагалось, что на каждый отдельный процессор будет добавлен свой
// элемент массива, который будет инкрементирован, но все ядра все равно
// шарили общий счетчик, поскольку размер cache линии = 128 байт,
// а размер счетчика - 4 байта. Соответсвенно, на одну cache-линию
// поместится много счетчиков, которые в момент изменения будут сбрасываться.
// И когда cache-линии будут изолированы дополнительной памятью - получилось
// нормальное разделение по процессорам, то есть ядра перестают друг другу мешать.
// То есть проблема заключалась в том, что мы хотели разделять данные, но из-за
// того, что они помещаются в cache-линию возникают проблемы.

// go test -bench=. perf_test.go

type MutexCounter struct {
	value int32
	mutex sync.Mutex
}

func (c *MutexCounter) Increment(int) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.value++
}

func (c *MutexCounter) Get() int32 {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return c.value
}

type AtomicCounter struct {
	value atomic.Int32
	// Добавим сюда некоторое выравнивание - [60] - некоторая пустая память
}

func (c *AtomicCounter) Increment(int) {
	c.value.Add(1)
}

func (c *AtomicCounter) Get() int32 {
	return c.value.Load()
}

// В силу того, что атомик выполняется на нескольких процессорах -
// можно пошардировать эту структуру на массив из 10 элементов.
type ShardedAtomicCounter struct {
	shards [10]AtomicCounter
}

// Будем инкрементировать значение внутри массива, которое является атомиком.
func (c *ShardedAtomicCounter) Increment(idx int) {
	c.shards[idx].value.Add(1)
}

// В Get-методе просто складываются все значения в одно.
func (c *ShardedAtomicCounter) Get() int32 {
	var value int32
	for idx := 0; idx < 10; idx++ {
		value += c.shards[idx].Get()
	}

	return value
}

// MutexCounter = 667.2 ns/op - использование mutex-а для синхронизации

// Код с использование атомиков выполняется в два раза быстрее = 260.2 ns/op

// Код с шардингом будет исполняться примерно за такое же время - 260.1 ns/op

// После добавления выравнивания в структуру AtomicCounter - код стал выполняться в разы быстрее -  46.29 ns/op

// Если вместо 60 байт добавить дополнительно памяти на 124 байта, то код ускорится еще больше - 33.75 (26.74) ns/op

func BenchmarkAtomicCounter(b *testing.B) {
	wg := sync.WaitGroup{}
	wg.Add(runtime.NumCPU())

	counter := MutexCounter{}
	// counter := AtomicCounter{} -  здесь смысл в том, что все ядра
	// шарят атомик, который будет подгружен в cache-линию.
	// Соответственно, если атомик модифицируется в рамках одного
	// процессора, то в остальных он будет сбрасываться. Следовательно,
	// чем больше ядер, тем сильнее они начинают друг другу мешать.
	for i := 0; i < runtime.NumCPU(); i++ {
		go func(idx int) {
			defer wg.Done()
			for j := 0; j < b.N; j++ {
				counter.Increment(idx)
			}
		}(i)
	}

	wg.Wait()
}
