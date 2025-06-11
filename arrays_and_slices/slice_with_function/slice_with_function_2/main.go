package main

import "fmt"

func main() {
	data := make([]int, 3, 6)

	fmt.Println("initial slice:", data) // [0, 0, 0]
	process(data)
	fmt.Println("after process:", data) // [0, 0, 0]

	// После расширения слайса получается так,
	// что, поскольку в слайсе есть место -
	// 5 попадет в слайс, но при этом длина исходного
	// массива не будет модифицирована, поскольку
	// после результата работы функции append
	// будет модифицирована длина в копии структуры
	// и как следствие при выведении этого слайса
	// ничего не увидим.

	// fmt.Println("after process:", data[:4]) // [0, 0, 0, 5]
}

func process(data []int) {
	data = append(data, 5)
}
