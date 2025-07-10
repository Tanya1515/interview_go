package main

import "fmt"

// Вывод:
// nil
// nil
// nil
// panic: error

// Происходит так по той причине, что вызов функции
// recover полезен только внутри функции defer.
// В противном случае, recover просто вернет nil и
// больше не будет ни на что влиять. Пскольку когда
// летит паника происходит stack unwiding - раскрутка стека.
// Когда раскручивается стек, панику можно обработать,
// а в альтернативных ситуациях - нет.

// Здесь вызывается печать recover() вначале и в конце,
// а по середине вызывается функция panic() со значением error.
func process2() {
	fmt.Println(recover())
	panic("error")
	fmt.Println(recover())
}

// Здесь вызывается печать recover() вначале и в конце,
// а по середине вызывается функция process2.
func process1() {
	fmt.Println(recover())
	process2()
	fmt.Println(recover())
}

// Здесь вызывается печать recover() вначале и в конце,
// а по середине вызывается функция process1.
func main() {
	fmt.Println(recover())
	process1()
	fmt.Println(recover())
}
