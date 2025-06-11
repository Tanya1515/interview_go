package main

import "fmt"

// Стоит помнить, что при передаче функции в слайс передается копия структуры,
// а именно указатель, длина и capacity слайса. Таким образом, если модифицировать
// данные в рамках существующей длины - они будут меняться, а если модифицировать
// данные за пределами len - они меняться не будут, поскольку будет меняться только
// копия структуры.

func main() {
	data := []int{1, 2, 3, 4} // len = 4, cap = 4

	fmt.Println("initial slice: ", data) // [1,2,3,4]
	process1(data)
	fmt.Println("after process1:", data) // [1,2,5,4]
	process2(data)
	fmt.Println("after process2:", data) // [1,2,5,4]
}

func process1(data []int) {
	data[2] = 5
}

func process2(data []int) {
	data = append(data, 6)

	// len: 5 cap:8 - слайс вырастет в размере и переедет в другую область памяти

	fmt.Println("len:", len(data), "cap:", cap(data))
}
