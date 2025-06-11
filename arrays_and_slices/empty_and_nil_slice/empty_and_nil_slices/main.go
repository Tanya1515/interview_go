package main

import (
	"fmt"
	"unsafe"
)

// len(data) == 0 - проверка на то, а является ли срез пустым
// data == nil - является ли срез нулевым
// unsafe.Sizeof(data) - размер слайса, стоит учитывать, что
// если срез нулевой, то он никуда и не указывает
// unsafe.SliceData(data) - куда указывает указатель

// Стоит учитывать, что срез может создаваться нулевым
// но не пустым - 3, 4 случаи. В этих случаях будут
// существовать указатели, которые будут указывать на некоторую
// область памяти.

func main() {
	var data []string
	fmt.Println("var data []string:") //

	// empty=true nil=true size=24 data=0x0
	fmt.Printf("\tempty=%t nil=%t size=%d data=%p\n", len(data) == 0, data == nil, unsafe.Sizeof(data), unsafe.SliceData(data))

	data = []string(nil)
	fmt.Println("data = []string(nil):") //

	// empty=true nil=true size=24 data=0x0
	fmt.Printf("\tempty=%t nil=%t size=%d data=%p\n", len(data) == 0, data == nil, unsafe.Sizeof(data), unsafe.SliceData(data))

	data = []string{}
	fmt.Println("data = []string{}:") //

	// empty=true nil=false size=24 data={{ адрес }}
	fmt.Printf("\tempty=%t nil=%t size=%d data=%p\n", len(data) == 0, data == nil, unsafe.Sizeof(data), unsafe.SliceData(data))

	data = make([]string, 0)
	fmt.Println("data = make([]string, 0):") //

	// empty=true nil=false size=24 data={{ адрес }}
	fmt.Printf("\tempty=%t nil=%t size=%d data=%p\n", len(data) == 0, data == nil, unsafe.Sizeof(data), unsafe.SliceData(data))

	empty := struct{}{}

	//
	fmt.Println("empty struct address:", unsafe.Pointer(&empty))
}
