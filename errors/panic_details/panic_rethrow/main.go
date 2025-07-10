package main

import (
	"fmt"
)

// Здесь представлена возможность перехвата паники
// и дальнейшего ее проброса в следующую функцию
// с defer-ом. В данном примере код корректно отработает
// и будет выведено:
// #2 recovered: error
// #1 recovered: error

func process2() {
	defer func() {
		e := recover()
		fmt.Println("#2 recovered:", e)
		panic(e)
	}()

	panic("error")
}

func process1() {
	defer func() {
		fmt.Println("#1 recovered:", recover())
	}()

	process2()
}

func main() {
	process1()
}
