package main

import (
	"fmt"
	"runtime"
)

// Функция main тоже считается отдельной горутиной.
func main() {
	fmt.Printf("Goroutines: %d\n", runtime.NumGoroutine()) // Goroutines: 1
}
