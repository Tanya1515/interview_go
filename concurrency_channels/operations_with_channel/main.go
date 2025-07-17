package main

import (
	"fmt"
)

// При записи в nil-канал будет deadlock
// (поскольку происходит блокировка на все время)
func writeToNilChannel() {
	var ch chan int
	ch <- 1
}

// Здесь происходит запись в закрытый канал. Будет паника.
func writeToClosedChannel() {
	ch := make(chan int, 2)
	close(ch)
	ch <- 20
}

// Создается буферезированный канал, в который
// записывается два значения. Далее считываем
// значение из канала и печатаем его. Затем
// закрываем канал и снова считываем из него значение,
// печатаем его. Код отработает корректно, причем будет
// напечатано: 10 true, 20 true, 0 false
func readFromChannel() {
	ch := make(chan int, 2)
	ch <- 10
	ch <- 20

	val, ok := <-ch
	fmt.Println(val, ok)

	close(ch)
	val, ok = <-ch
	fmt.Println(val, ok)

	val, ok = <-ch
	fmt.Println(val, ok)
}

// Создается два канала, а затем в оба канала записываются данные.
// После чего в select-е считываются данные и выводятся на экран.
// Порядок выбора select-а производится рандомно.
func readAnyChannels() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		ch1 <- 100
	}()

	go func() {
		ch2 <- 200
	}()

	select {
	case val1 := <-ch1:
		fmt.Println(val1)
	case val2 := <-ch2:
		fmt.Println(val2)
	}
}

// Попытка прочитать из nil-канала.
// Произойдет вечная блокировка горутины.
func readFromNilChannel() {
	var ch chan int
	<-ch
}

// Итерация по nil-каналу - будет опять же блокировка,
// поскольку будем читать из канала, пока он не закрыт.
func rangeNilChannel() {
	var ch chan int
	for range ch {

	}
}

// Если закрыть nil-канал, то будет panic-а.
func closeNilChannel() {
	var ch chan int
	close(ch)
}

// Если закрыть канал несколько раз - будет panic.
func closeChannelAnyTimes() {
	ch := make(chan int)
	close(ch)
	close(ch)
}

// Каналы можно сравнивать между собой, причем будет напечатано:
// false true. В силу того, что каналы - это указатели,
// то если указатели равны, то и каналы равны.
func compareChannels() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	equal1 := ch1 == ch2
	equal2 := ch1 == ch1

	fmt.Println(equal1)
	fmt.Println(equal2)
}

func main() {
}
