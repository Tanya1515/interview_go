package main

import "unsafe"

// Данный пример доказывает, что строку изменять можно, но
// только в зависимости от того, где она располагается.

func action() {}

func main() {
	var str = "go"
	newStr := str + "-go"

	// Получем указатель на массив, который используется
	// для строки
	strData := unsafe.StringData(str)
	newStrData := unsafe.StringData(newStr)

	// Здесь выводятся адреса всех переменных и функции action
	// Если запустить этот код, то будет видно, что фадрес
	// функции и strData примерно одинаковые, это означает,
	// что они располагаются в одинаковом сегменте памяти
	// - read only сегмент, в котором хранится спкомпилированный
	// код программы.
	println("action:", action)
	println("strData:", strData)
	println("newStrData:", newStrData)

	slice := unsafe.Slice(strData, len(str))
	newSlice := unsafe.Slice(newStrData, len(newStr))

	newSlice[0] = 'G'
	println("newStr:", newStr)

	slice[0] = 'G' // здесь будет паника
	println("str:", str)
}
