package main

import (
	"fmt"
	"runtime"
	"sync"
)

/*

	Livelock - ситуация, в которой система не застревает,
	а занимается бесполезной работой. Ее состояние постоянно
	меняется, но при этом она не производит никакой полезной
	работы.

	В коде ниже в разном порядке захватываются мьютексы, а затем
	каждая из горутин пытается захватить чужой мьютекс при помощи
	операции TryLock, причем если не получается горутина начинает
	исполнять пустую работу.

	Livelock чаще всего встречаются в lockfree структурах данных.

*/

var mutex1 sync.Mutex
var mutex2 sync.Mutex

func goroutine1() {
	mutex1.Lock()

	runtime.Gosched()
	for !mutex2.TryLock() {
		// active waiting
	}

	mutex2.Unlock()
	mutex1.Unlock()

	fmt.Println("goroutine1 finished")
}

func goroutine2() {
	mutex2.Lock()

	runtime.Gosched()

	// TryLock постоянно пытается захватить мьютекс, при этом
	// если не удается - то он просто возвращает false.
	// Если мьютекс получается захватить - то он захватывается
	// и возвращается true.

	for !mutex1.TryLock() {
		// active waiting
	}

	mutex1.Unlock()
	mutex2.Unlock()

	fmt.Println("goroutine2 finished")
}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		goroutine1()
	}()

	go func() {
		defer wg.Done()
		goroutine2()
	}()

	wg.Wait()
}
