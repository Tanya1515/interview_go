package main

import (
	"fmt"
	"sync"
)

// Стоит помнить, что лучше всего выполнять минимальное
// количество действий в критической секции.

var mutex sync.Mutex
var cache map[string]string

func doSomething() {
	mutex.Lock()
	item := cache["key"]
	fmt.Println(item)
	mutex.Unlock()

	// не имеет смысла делать операцию ввода выод
	// под блокировкой, поскольку в данном случае
	// создается своя переменная item для каждой горутины.
	// Это же касается аналогичных походов в сеть и так далее.
	// fmt.Println(item)
}
