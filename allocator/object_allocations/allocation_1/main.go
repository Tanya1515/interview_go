package main

// -l = disable inlining
// -m = print optimization decisions
// go build -gcflags '-l -m'

// В данном варианте объект будет проаллоцирован на стеке,
// а в main вернется копия значения.

func getResult() int {
	result := 200
	return result
}

func main() {
	_ = getResult()
}
