package main

import "fmt"

func main() {
	// Есть целочисленное значение, которое
	// прячется за пустым интерфейсом
	var value int = 100
	var i interface{} = value

	// Далее производится попытка вывести type assertion -
	// то есть проверить, а скрывается ли тип int
	// за пустым интерфейсом и удастся ли привести его к int.
	converted1, ok1 := i.(int)
	if ok1 {
		fmt.Println("converted1 int:", converted1) // отработает корректно
	}

	// Аналогично здесь, но уже к float32.
	converted2, ok2 := i.(float32)
	if ok2 {
		// будет паника, поскольку в пустом интерфейсе
		// совершенно другой тип данных
		fmt.Println("converted2 float32:", converted2)
	}

	converted3 := i.(int)
	fmt.Println("converted3 int:", converted3) // отработает корректно
	converted4 := i.(float32)
	fmt.Println("converted4 float32:", converted4) // будет паника
}
