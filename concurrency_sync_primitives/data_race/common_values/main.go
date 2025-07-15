package main

import "sync"

type Data struct {
	X int
	Y int
}

/*

	В данном коде нет data race, поскольку в первой и во второй горутинах
	происходит обращение к различным участкам памяти. Слайс и структура - это
	аггрегирующие участки памяти, поэтому поля X и Y - это различные участки
	памяти, равно как и слайс.

*/

func main() {
	var data Data
	values := make([]int, 2)

	wg := sync.WaitGroup{}

	wg.Add(2)

	go func() {
		defer wg.Done()

		data.X = 5
		values[0] = 5
	}()

	go func() {
		defer wg.Done()

		data.Y = 10
		values[1] = 10
	}()

	wg.Wait()
}
