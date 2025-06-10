package main

import (
	"fmt"
	"math"
)

func main() {
	// восьмибитная знаковая целочисленная переменная, которой присваивается
	// максимальное значение, которое может вместить в себя эта переменная
	var signed int8 = math.MaxInt8
	signed++

	// беззнаковая восьмибитная переменная и здесь аналогично
	var unsigned uint8 = math.MaxUint8
	unsigned++

	// при переполнении выодится самый маленький элемент для этого типа данных
	fmt.Println(signed)   // -128
	fmt.Println(unsigned) // 0

	// в этом варианте компилятор поймет, что здесь происходит переполнение
	// и не даст увеличить переменную

	// var signed int8 = math.MaxInt8 + 1     -> compilation error
	// var unsigned uint8 = math.MaxUint8 + 1 -> compilation error
}
