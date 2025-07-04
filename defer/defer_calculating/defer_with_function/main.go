package main

import "fmt"

func get() string {
	fmt.Println("1")
	return ""
}

func handle(string) {
	fmt.Println("3")
}

func process() {
	defer handle(get())
	fmt.Println("2")
}

func main() {
	process()
}

/* 1 2 3

Сначала будет вычислен аргумент функции handle, то есть get-функция,
которая напечатает 1 и вернет пустую строку. Затем будет выполнен Println
функции process, по завершении которой будет выполнена функция handle,
то есть напечатается 3.

*/
