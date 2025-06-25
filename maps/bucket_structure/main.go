package main

import (
	"fmt"
	"unsafe"
)

// Бакеты в golang внутри себя хранят массив ключей
// и массив значений, поскольку так этот бакет занимает
// гораздо меньше места, чем если бы это был массив структур.

type bucketV1 struct {
	key   int8
	value int64
}

type bucketV2 struct {
	keys   [8]int8
	values [8]int64
}

func main() {
	fmt.Println("bucketsV1:", unsafe.Sizeof([8]bucketV1{})) // 128
	fmt.Println("bucketsV2:", unsafe.Sizeof(bucketV2{}))    // 72
}
