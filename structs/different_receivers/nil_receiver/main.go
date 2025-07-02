package main

import "fmt"

// Стоит учитывать, что можно обращаться
// к nil-овому receiver-у, при условии, что в самой функции
// не будет происходить разыменования receiver-а,
// а также не будет происходить обращения к полям nil-ового receiver-а.

// Однако, возникает вопрос: почему не происходит разыменования receiver-a
// при вызове его метода? Поскольку сам receiver - это
// синтаксический сахар, под капотом поля receiver-а передаются
// внутрь функции как аргументы (в частности в качестве первого аргумента).

type Obect struct{}

func (o *Obect) Print() {
	if o == nil {
		fmt.Println("nil")
	} else {
		fmt.Println("not nil")
	}
}

func main() {
	var object *Obect
	object.Print()
}
