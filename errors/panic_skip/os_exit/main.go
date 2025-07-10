package main

import (
	"fmt"
	"os"
)

// Здесь же при вызове os.Exit defer работать не будет.
func process() {
	defer func() {
		recover()
	}()

	fmt.Println("V2: open file")
	defer fmt.Println("V2: close file")

	os.Exit(1)
}

func main() {
	process()
}
