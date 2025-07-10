package main

import (
	"fmt"
	"runtime"
)

// Даже просто при завершении работы горутины, без обработки паники
// defer тоже отработает.

func process() {
	defer func() {
		recover()
	}()

	fmt.Println("V2: open file")
	defer fmt.Println("V2: close file")

	runtime.Goexit()
}

func main() {
	process()
}
