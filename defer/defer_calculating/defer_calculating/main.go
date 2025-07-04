package main

import "fmt"

// Need to show solution

const (
	StatusOk    = "ok"
	StatusError = "error"
)

// Вычисление аргументов функции defer происходит сразу же,
// а не после выхода из окружающей функции (следовательно,
// golang откладывает выполнение функции с тем значением,
// которое было передано при вызове defer.)

func notify(status string) {
	fmt.Println(status) // в качестве вывод здесь будет пустая строка.
}

func process() {
	var status string

	/*
		Однако, если замкнуть значение status
		 в отложенном вызове, то в функции
		 окажется последнее вызванное значение

			defer func(status) {
				notify(status)
			}()

			В этом случае будет напечатано "error",
			поскольку мы будем ссылаться на итоговое значение

			Еще один способ - использовать указатели:

			defer func(s *string) {
				notify(*s)
			}(&status)

	*/

	// Здесь откладывается вызов анонимной функции,
	// в рамках которой вызывается функция notify(s)
	defer func(s string) {
		notify(s)
	}(status) // при вызове defer значение status копируется.

	// processing..
	status = StatusError
}

func main() {
	process()
}
