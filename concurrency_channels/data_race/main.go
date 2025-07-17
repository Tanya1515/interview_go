package main

import "sync"

// go run -race main.go

// Канал представляет собой указатель на структуру hchan.
// То есть это адрес [0xA0] и получается, что по этому
// адресу из множества горутин записывается некоторое
// значение. По этой причине возникает data race при
// выполнении кода.
var buffer chan int

func main() {
	wg := sync.WaitGroup{}
	// Есть WaitGroup-а, которая ожидает 100 горутин.
	wg.Add(100)

	for i := 0; i < 100; i++ {
		// В каждой горутине создается канал для передачи
		// элементов типа int.
		go func() {
			defer wg.Done()
			buffer = make(chan int)
		}()
	}

	wg.Wait()
}
