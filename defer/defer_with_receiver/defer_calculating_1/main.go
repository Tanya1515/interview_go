package main

import "fmt"

type data1 struct {
	value int
}

// здесь ресивер используется по значению
func (d data1) print() {
	fmt.Println("data1", d.value)
}

type data2 struct {
	value int
}

// Здесь ресивер используется по указателю
func (d *data2) print() {
	fmt.Println("data2", d.value)
}

func main() {
	// Создается объект структуры data1
	// и откладывается вызов функции print
	d1 := data1{}
	// Здесь когда отклаывается выполнение функции
	// происходит копирование ресивера, который
	// по своей сути явялетс первым аргументом в функции
	defer d1.print()

	// Создается объект структуры data2
	// и откладывается вызов функции print
	d2 := data2{}
	defer d2.print()

	// Далее у первого и второго объектов
	// меняются значения
	d1.value = 100
	d2.value = 200
}

// Вывод:
// data2 200
// data1 0
