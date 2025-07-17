package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var data map[string]string
var initialized atomic.Bool

// В коде ниже нет data race, однако, все равно мапа может
// быть проинициализирована несколько раз
func initialize() {
	if !initialized.Load() {
		/*

			Поскольку между этим строками может встроиться другая горутина.
			По этой причине необходимо Read/Write-Modify операция.
			Эта операция - CompareAndSwap.

			initialized.CompareAndSwap(false, true) - первым аргументом передается
			старое значение, вторым аргументом передается новое значение. При успешном
			выполнеии операции будет возвращено true. Эта операция выполняется атомарно и
			ровно один раз.

			На псевдокоде CAS устроен так:

				if *addr == old {
					*addr = new
					return true
				}

				return false

		*/
		initialized.Store(true)
		data = make(map[string]string)
		fmt.Println("initialized")
	}
}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1000)

	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()
			initialize()
		}()
	}

	wg.Wait()
}
