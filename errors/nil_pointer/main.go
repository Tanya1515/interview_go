package main

import "fmt"

// В момент разыменования nil-pointer-а возникает ошибка.
// В функции dererence передается указатель на nil, и
// паника, которая будет вызвана при попытке разыменования
// nil-указателя, будет обработана.

func dererence(pointer *int) {
	defer func() {
		fmt.Println("recovered:", recover())
	}()

	_ = *pointer
}

func main() {
	dererence(nil)
	// Будет выведено обработка паники и Hello!
	fmt.Println("Hello!")
}
