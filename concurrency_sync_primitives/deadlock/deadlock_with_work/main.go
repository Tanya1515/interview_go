package main

import (
	"log"
	"sync"
	"time"
)

var resource1 int
var resource2 int

func normalizeResources(lhs, rhs *sync.Mutex) {
	lhs.Lock()
	rhs.Lock()

	// normalization

	rhs.Unlock()
	lhs.Unlock()
}

// Стоит отметить, что runtime golang не всегда сможет определить deadlock.
// Причем здесь как только будет deadlock из-за некорректного порядка передачи
// mutex-ов в функции в разных горутинах, просто запустится функция, которая будет
// тикать, но самой ошибки не будет (будет работать wg.Wait()).

func main() {
	var mutex1 sync.Mutex
	var mutex2 sync.Mutex

	wg := sync.WaitGroup{}
	wg.Add(1000)

	for i := 0; i < 500; i++ {
		go func() {
			defer wg.Done()
			normalizeResources(&mutex1, &mutex2)
		}()
	}

	for i := 0; i < 500; i++ {
		go func() {
			defer wg.Done()
			normalizeResources(&mutex2, &mutex1)
		}()
	}

	time.Sleep(time.Millisecond * 100)

	go func() {
		for {
			time.Sleep(time.Second)
			log.Println("tick")
		}
	}()

	wg.Wait()
}
