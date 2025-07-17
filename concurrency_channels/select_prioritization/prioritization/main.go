package main

import (
	"fmt"
	"time"
)

func producer(ch chan<- int) {
	for {
		ch <- 1
		time.Sleep(time.Second)
	}
}

func main() {
	// Есть несколько каналов, причем первый из
	// них более приоритетный, чем второй.
	ch1 := make(chan int) // more prioritized
	ch2 := make(chan int)

	// создаем две горутины, которые раз
	// в секунду будут писать в канал.
	go producer(ch1)
	go producer(ch2)

	for {
		// Здесь если значение в канале ch1 и канале ch2 будет
		// находится значение, то выбор ветки будет абсолютно рандомным.

		/*
			Второй способ - создать еще один неблокирующий select
			до основного select-а. Таким образом, уменьшается вероятность
			того, что в обоих каналах одновременно наступят два события.

			select {
			case value := <- ch1:
				fmt.Println(value)
				return
			default:
			}
		*/
		select {
		case value := <-ch1:
			fmt.Println(value)
			return
		case value := <-ch2:
			/*
				Первый способ приоритезации - создать вложенный select,
				причем select неблокирующий, то есть с условием default.

				select {
				case value := <-ch1:
					fmt.Println(value)
					return
				default:
				}
			*/
			fmt.Println(value)
		}
	}
}
