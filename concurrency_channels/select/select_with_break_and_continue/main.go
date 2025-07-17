package main

import "fmt"

func main() {

	data := make(chan int)
	// В этой горутине в канал записывается 4 значения,
	// а затем канал закрывается.
	go func() {
		for i := 1; i <= 4; i++ {
			data <- i
		}
		close(data)
	}()

	// В рамках цикла производится чтение из канала,
	// при получении элемента равного 3 - выходим из
	// бесконечного цикла. Если канал закрыт - выходим из функции.
	for {
		value := 0
		opened := true

		select {
		case value, opened = <-data:
			if value == 2 {
				// Относится к циклу, то есть при выполнении
				// этой операции - мы передем на следующую итерацию
				continue
			} else if value == 3 {
				// Отработает на select
				break
			}

			if !opened {
				return
			}
		}

		fmt.Println(value)
	}
}
