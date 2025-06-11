package main

import (
	"fmt"
	"runtime"
)

// Пусть есть некоторая структура, полями
// которой является срез байт.
//

type Data struct {
	values []byte
}

func printAllocs() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%d MB\n", m.Alloc/1024/1024)
}

func GetData() []Data {
	// здесь происходит аллокация памяти под срез,
	// элементами которого являются структуры типа Data
	data := make([]Data, 1000) // ~24KB

	// Далее в цикле производится аллокация
	// для каждого поля values выделяется массив
	// из 1<<20 байт.
	for i := 0; i < len(data); i++ {
		data[i] = Data{
			values: make([]byte, 1<<20),
		}
	}

	//clear(data[2:])
	return data[:2]
}

// Каждый элемент слайса указывает на достаточно тяжелый объект
// Причем опять возвращается из функции только часть среза.
// Таким образом, чтобы не держать ту часть памяти, которая
// рассчитана под остальные элементы - ее можно почистить.

func main() {
	data := GetData()

	printAllocs()
	runtime.GC()
	printAllocs()

	runtime.KeepAlive(data)
}
