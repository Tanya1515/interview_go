package main

import "fmt"

// -l = disable inlining
// -m = print optimization decisions
// go build -gcflags '-l -m'

// Есть некоторая функция, которая
// в качестве параметра принимает пустой интерфейс
// и печатает значение
func printValue(v interface{}) {
	fmt.Println(v)
	//_, _ = v.(int)
}

// Однако, если функцию fmt.Println() заменить на println,
// то num1, str1 будут выделены на стеке. Так происходит,
// поскольку компилятор не может доказать, что fmt.Println()
// не будет использовать v за пределами функции, поэтому
// они пробрасываются через heap.

// Стоит помнить, что escape analyze использует такие же правила
// как для объектов типа int, так и для пустых интерфейсов.

func main() {
	var num1 int = 10         // heap
	var str1 string = "Hello" // heap

	printValue(num1)
	printValue(str1)

	var num2 int = 10         // stack
	var str2 string = "Hello" // stack

	// здесь значения num2 и str2 присваиваются пустому интерфейсу.
	var i interface{}
	i = num2
	i = str2
	_ = i
}
