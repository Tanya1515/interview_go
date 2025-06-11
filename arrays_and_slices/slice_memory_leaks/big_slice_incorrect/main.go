package main

import (
	"fmt"
	"runtime"
)

func printAllocs() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%d MB\n", m.Alloc/1024/1024)
}

// FindData аллоцирует слайс большого размера,
// после чего проходится по этому слайсу и возвращает
// слайс, размер которого составляет всего лишь 20 элементов.
//
// Затем происходит вызов сборщика мусора, однако, слайс data не будет очищен
// поскольку на него ссылается другой слайс, в котором всего лишь 20 элементов
//
// Таким образом, возникает утечка памяти - используется только 20 элементов
// которые ссылаются на слайс большего размера.
//
// Чтобы избежать такой утечки памяти, можно скопировать слайс:
// то есть вызвать функцию copy(partData, data)
// и вернуть полную копию (deep copy) от среза слайсов

func FindData(filename string) []byte {
	data := make([]byte, 1<<30) // for example read from file

	for i := 0; i < len(data)-1; i++ {
		if data[i] == 0x00 && data[i+1] == 0x00 {
			return data[i : i+20]
		}
	}

	return nil
}

func main() {
	data := FindData("data.bin")
	_ = data

	printAllocs()
	runtime.GC()
	printAllocs()

	runtime.KeepAlive(data)
}
