package main

import "fmt"

// Вывод:
// recovered: 3

// То есть в данном случае будет реализована подмена паники.

// Важно понимать: в любой момент времени функция может быть
// ассоциирована не более, чем с одной необработанной паникой.

func main() {
	defer func() {
		fmt.Println("recovered:", recover())
	}()

	defer panic(3) // replaced panic(3)
	defer panic(2) // replaced panic(3)
	defer panic(1) // replaced panic(3)
	// в панике выбрасывается 0
	panic(0)
}
