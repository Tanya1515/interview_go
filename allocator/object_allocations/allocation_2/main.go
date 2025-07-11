package main

// -l = disable inlining
// -m = print optimization decisions
// go build -gcflags '-l -m'

// Здесь возвращается указатель на значение,
// которое было создано в функции getResult.
// Поэтому переменная result, поэтому переменная result
// будет аллоцирована в куче. Поскольку в
// ином случае будем ссылаться на участок памяти,
// которого не существует. То есть будет проблема nil pointer dereference.

func getResult() *int {
	result := 200
	return &result
}

func main() {
	_ = getResult()
}
