package main

// -l = disable inlining
// -m = print optimization decisions
// go build -gcflags '-l -m'

func createPointer() *int {
	// heap, поскольку это значение используется за пределами функции createPointer
	value2 := new(int)
	return value2
}

func main() {
	// stack, поскольку за пределами функции никто не будет ссылаться на эту переменную
	value1 := new(int)
	_ = value1

	value2 := createPointer() // heap
	_ = value2
}
