package main

import "fmt"

/*
Стоит помнить, что defer выполняются в обратном порядке,
то есть в режиме стека.
*/
func process1() {
	defer fmt.Println(3)
	defer fmt.Println(2)
	defer fmt.Println(1)
}

func process2() {
	// Сначала откладывается выполнение первой функции
	defer func() {
		defer fmt.Println(4)
		defer fmt.Println(3)
	}()

	// Откладываем выполнение второй функции
	defer func() {
		defer fmt.Println(2)
		defer fmt.Println(1)
	}()
}

// После завершения начинает работу сначала вторая функция, только потом первая

func main() {
	process1() // 1 2 3
	process2() // 1 2 3 4
}
