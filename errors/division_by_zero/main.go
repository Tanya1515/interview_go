package main

import "fmt"

/*
При делении на 0 в golang происходит паника,
которую можно обработать при помощи recover

То есть при выполнении этой функции будет выведено:

	recovered: runtime error integer divide by zero

Однако, если поменять тип данных с int на float32,
то никакая паника обработана не будет.
Поскольку при делении на 0, будет выведена бесконечность.
Это связано с тем, как кодируются целочисленный
тип данных и числа с плавающей точкой.
*/
func divide(lhs, rhs int) int {
	defer func() {
		fmt.Println("recovered:", recover())
	}()

	return lhs / rhs
}

func main() {
	_ = divide(1000, 0)
}
