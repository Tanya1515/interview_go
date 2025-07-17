package main

import "fmt"

func main() {
	ch1 := make(chan struct{}, 1)
	ch2 := make(chan struct{}, 1)

	// закрытие канала производится с целью, чтобы эмулировать
	// постоянное чтение из канала, поскольку из закрытого
	// канала всегда считывается ZeroValue.
	close(ch1)
	close(ch2)

	ch1Value := 0.0
	ch2Value := 0.0

	for i := 0; i < 100000; i++ {
		select {
		// Здесь дублируется условие case-а, чтобы увеличить
		// вероятность выбора данных из ch1.
		case <-ch1:
			ch1Value++
		case <-ch1:
			ch1Value++
		case <-ch2:
			ch2Value++
		}
	}

	// Действительно, если выполнить код, то счетчик для ch1 будет
	// в несколько раз больше, чем счетчик для ch2.
	fmt.Println(ch1Value / ch2Value)
}
