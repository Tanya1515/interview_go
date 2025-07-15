package main

import "sync"

/*

	Deadlock - ситуация в многозадачной среде, при которой
	несколько потоков (горутин) находятся в состоянии ожидания
	ресурсов, занятых друг другом, и ни один из них не может
	продолжать свое исполнение. То есть обе эти горутины
	были заблокированы и перешли в состояние waiting.

*/

var resource1 int
var resource2 int

func normalizeResources(lhs, rhs *sync.Mutex) {
	lhs.Lock()
	rhs.Lock()

	// normalization

	rhs.Unlock()
	lhs.Unlock()
}

// Проблема кода ниже заключается в том,
// что mutex-ы передаются в функции в различном порядке.
// Из-за этого может возникать deadlock. В качестве решения
// необходимо согласовать порядок блокировок.

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

	wg.Wait()
}
