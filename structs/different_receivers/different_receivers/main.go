package main

import "fmt"

// Вывод при работе с данными структурами будет разным,
// поскольку в первом случае receiver копируется,
// а во втором - передается по указателю.

// Receiver должен быть указателем:
// 1) Если метод должен изменить receiver
// 2) Если в receiver есть поле, которое нельзя копировать
// (например, штип из пакета sync)

// Receiver следует сделать указателем:
// 1) Если receiver -это большой объект (чтобы каждый раз
// его не копировать, при этом размер обхъекта необходимо бенчмаркать)

// Receiver должен быть значением:
// 1) Если нужно обеспечить неизменяемость receiver.

// Reciever следует сделать значением:
// 1) Если receiver является базовым типом.
// 2) Если изменяемые поля не являются частью receiver,
// а находятся частью другой структуры. (пример в recursive_receiver)

type customer1 struct {
	balance int
}

func (c customer1) add(value int) {
	c.balance += value
}

type customer2 struct {
	balance int
}

func (c *customer2) add(value int) {
	c.balance += value
}

func main() {
	c1 := customer1{}
	c1.add(100)

	c2 := customer2{}
	c2.add(100)

	fmt.Println(c1) // {0}
	fmt.Println(c2) // {100}
}
