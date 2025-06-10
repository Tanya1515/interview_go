package main

import "fmt"

// Стоит учитывать, что при передаче указателя в функцию
// он копируется. То есть создается некоторый temp, который
// будет указывать на другую переменную. Таким образом, необходимо
// передавать указатель на указатель.

func process(temp **int32) {
	// здесь важно, чтобы указатель pointer
	// теперь указывал на value2.
	var value2 int32 = 200
	// здесь разыменование проиходит один раз, чтобы достучаться до указателя.
	*temp = &value2
}

func main() {
	var value1 int32 = 100
	// берем указатель на это целочисленное значение.
	pointer := &value1

	fmt.Println("value:", *pointer)
	fmt.Println("address:", pointer)

	// передаем указатель в функцию
	process(&pointer)

	fmt.Println("value", *pointer)
	fmt.Println("address:", pointer)
}
