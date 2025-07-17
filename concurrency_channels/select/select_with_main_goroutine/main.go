package main

import "runtime"

// при вызове двух горутин, которые будут выполнять этот код
// - горутины будут переключаться между собой при помощи планировщика.
func doSomething() {
	for {
		runtime.Gosched()
	}
}

func main() {
	go doSomething()
	go doSomething()
	// горутина main будет заблокирована.
	select {}
}
