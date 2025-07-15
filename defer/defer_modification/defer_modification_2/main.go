package main

import "fmt"

func Modify(value int) int {
	var result int
	defer func() {
		result += value
	}()

	return value + value // 10
}

func main() {
	fmt.Println(Modify(5)) // 10
}
