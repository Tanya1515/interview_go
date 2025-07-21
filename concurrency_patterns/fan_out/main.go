package main

import (
	"fmt"
	"sync"
	"time"
)

// Паттерн Fan Out - паттерн, противоположный Fan In, 
// который разделяет один канал на несколько.  
// Алгоритм: 
// 1) Создали результирующий канал 
// 2) Вернули резульирующий канал 
// 3) Асинхронно начинаем обрабатывать данные


func SplitChannel(inputCh <-chan int, n int) []chan int {
	if n <= 0 {
		n = 1
	}

	// Создается срез результирующих каналов 
	outputCh := make([]chan int, n)
	for i := 0; i < n; i++ {
		outputCh[i] = make(chan int)
	}

	// Асинхронно начинаем процессить 
	go func() {
		idx := 0
		// Получаем итеративно значения из входного канала
		for value := range inputCh {
			// По индексу выбираем канал, в который будут записаны данные.
			// Для неблокирующей записи можно сделать буферезированный канал
			outputCh[idx] <- value 
			// Инкрементируем индекс 
			idx = (idx + 1) % n
		}

		// Проходимся по каждому из каналов и закрываем его. 
		for _, ch := range outputCh {
			close(ch)
		}
	}()

	// Возвращаем срез результирующих каналов 
	return outputCh
}

func main() {
	// создается канал
	ch := make(chan int)

	// В отдельной горутине производится запись 
	// данных в канал и с таймаутом на сон.
	go func() {
		defer close(ch)
		for i := 0; i < 10; i++ {
			ch <- i
			time.Sleep(100 * time.Millisecond)
		}
	}()

	// Вызывается функция, которая разделяет один канал на два разных, 
	// в качестве результата возвращается срез, состоящий из двух каналов.
	channels := SplitChannel(ch, 2)

	wg := sync.WaitGroup{}
	wg.Add(2)

	// Создаются отдельные горутины, каждая из которых читает из своего канала.
	go func() {
		defer wg.Done()
		for value := range channels[0] {
			fmt.Println("ch1: ", value)
		}
	}()

	go func() {
		defer wg.Done()
		for value := range channels[1] {
			fmt.Println("ch2: ", value)
		}
	}()

	wg.Wait()
}
