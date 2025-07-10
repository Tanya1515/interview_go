package main

import "fmt"

// В Golang некоторые ошибки, такие как переполнение стека,
// нехватка памяти, не подлежат восстановлению. Как только
// они произойдут -  программа выйдет из строя.

// GOGC=off go run main.go - запуск задачи без сборщика мусора

var data []byte

func allocate() {
	defer func() {
		fmt.Println(recover())
	}()

	count := 0

	// Здесь в бесконечном цикле будет постоянно выделяться по 10 Mb
	// памяти, что в конечном итоге должно вызвать OOM. Однако,
	// никакого рекаверинга не будет - просто будет прислан сигнал killed.

	for {
		data = make([]byte, 1<<30)
		for idx := 0; idx < 1<<30; idx += 4096 {
			data[idx] = 100
		}

		fmt.Println("allocated GB:", count)
		count++
	}
}

func main() {
	allocate()
}
