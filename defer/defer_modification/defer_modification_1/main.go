package main

import "fmt"

/*

Стоит отметить, что благодаря defer можно менять возвращаемое значение.
Однако, это раюотает только в случае, если в функции используется
именованное возвращаемое значение.

Если же тело функции Modify будет выглядеть следующим образом:

func Modify(value int) int {
	var result int
	defer func() {
		result += value
	}()

	return value + value
}
 то возвращаемое значение будет равно 10.

*/

func Modify(value int) (result int) {
	defer func() {
		result += value // 15
	}()

	return value + value // 10
}

func main() {
	fmt.Println(Modify(5))
}
