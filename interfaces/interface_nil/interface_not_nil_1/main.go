package main

import "fmt"

type MyError struct{}

func (e *MyError) Error() string {
	return "error"
}

func main() {
	// Объявляется указатель на тип MyError, который равен nil
	// тут это прописано явно, но даже если не присваивать
	// значение nil, это и так понятно.
	var pointer *MyError = nil

	// присваиваем интерфейсу error nil. MyError соответсвует
	// интерфейсу error, поскольку у интерфейса error есть
	// метод Error, который возвращает строку.
	var err error = pointer

	// В данном случае будет вывод false, поскольку интерфейс -
	// это указатель на itab и указатель на data. В data лежит nil,
	// а в itab не равен nil, поскольку там лежит указатель
	// на тип данных и набор методов, которые реализует этот тип данных.
	fmt.Println("nil:", err == nil)
}
