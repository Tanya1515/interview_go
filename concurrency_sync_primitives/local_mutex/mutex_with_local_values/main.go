package main

import (
	"log"
	"sync"
)

func main() {
	mutex := sync.Mutex{}
	wg := sync.WaitGroup{}
	wg.Add(1000)

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()

			value := 0
			for j := 0; j < 10; j++ {

				// здесь нет необходимости в мьютексе,
				// поскольку у каждой горутины создается
				// своя локальная переменная value.

				mutex.Lock()
				value++
				mutex.Unlock()
			}

			log.Println(value)
		}()
	}

	wg.Wait()
}
