package main

import "fmt"

func main() {
	data := [...]int{1, 2, 3}

	// Стоит отметить, что когда массив передается в функцию/range
	// то он копирует все элементы, а это долго и затратно по памяти
	
	for value := range data { // copy of array
		fmt.Println(value)
	}

	// Поэтому вместо тупого копирования лучше в range передавать указатель
	// на массив или передавать слайс, поскольку при передаче слайса будет
	// копироваться не все данные а только его структура - указатель, длина слайса
	// и capacity. 

	for value := range &data { // not a copy of array
		fmt.Println(value)
	}

	for value := range data[:] { // not a copy of array
		fmt.Println(value)
	}
}
