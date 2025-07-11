package main

// -l = disable inlining
// -m = print optimization decisions
// go build -gcflags '-l -m'

// Однако, не всегда, не всегда если использовать указатели,
// то переменная будет аллоцирована в куче. В задаче ниже
// в функции main создается переменная number, указатель
// на нее передается в функцию getResult. Внутри функции
// указатель разыменовывается и складывается с 200. Затем
// возвращается копия значения result.  Number будет аллоцирована
// на стеке, поскольку мы пробрасываем значение вверх по стеку.
// То есть когда функция getResult завершится - number будет
// существовать в области видимости функции main.

func getResult(number *int) int {
	result := *number + 200
	return result
}

func main() {
	number := 100
	_ = getResult(&number)
}
