package main

import "fmt"

func main() {
	data := map[int]int{1: 1, 2: 2, 3: 3}
	for _, value := range data {
		value = 1000
		_ = value
	}

	// Значения внутри мапы не поменяются, поскольку в golang
	// все копируется, а это означает, что в wbrk c range
	// будет передана копия. Но если использовать значения
	// по ключу - то, действительно, мапа меняться будет.
	fmt.Println(data)
}
