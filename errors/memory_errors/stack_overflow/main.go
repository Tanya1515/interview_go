package main

import "fmt"

/*
Функция, в рамках которой на каждом шаге рекурсии
выделяется 10 Mb. Память выделяется на стеке.
*/
func recursion() {
	var data = [10 * 1024 * 1024]int8{}
	_ = data

	recursion()
}

func main() {
	// Если вдруг произойдет паника, то это необходимо
	// зарекавэрить. Но обработать эту панику - переполнение
	// стека не получится.

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main:", r)
		}
	}()

	recursion()
}
