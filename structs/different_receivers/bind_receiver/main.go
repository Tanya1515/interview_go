package main

import "fmt"

// В этой программе представлены два типа receiver-а:
// Value1 - в качестве значения,
// Value2 - в качестве указателя.

type Data struct {
	value int
}

// Здесь происходит копирование receiver-а
func (d Data) Value1() int {
	return d.value
}

// Здесь в любом случае передается указатель.
func (d *Data) Value2() int {
	return d.value
}

func main() {
	data := Data{100}
	pointer := &data

	value1ByData := data.Value1       // функция
	value1ByPointer := pointer.Value1 // функция

	value2ByData := data.Value2       // функция
	value2ByPointer := pointer.Value2 // функция

	data.value = 200

	fmt.Println("value1ByData:", value1ByData())       // 100
	fmt.Println("value1ByPointer:", value1ByPointer()) // 100
	fmt.Println("value2ByData:", value2ByData())       // 200
	fmt.Println("value2ByPointer:", value2ByPointer()) // 200
}
