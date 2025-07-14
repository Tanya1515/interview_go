package main

import (
	"fmt"
	"time"
)

// Стоит отметить, что нельзя обработать панику, вызванную в другой горутине.
// Поскольку при возникновении паники возникает stack unwiding - расркутка стека.
// При этом у каждой горутины есть свой стек, поэтому из другой горутины нельзя поймать
// раскрутку стека другой горутины.

func process() {
	defer func() {
		v := recover()
		fmt.Println("recovered:", v)
	}()

	go func() {
		panic("error")
	}()

	time.Sleep(time.Second)
}

func main() {
	process()
}
