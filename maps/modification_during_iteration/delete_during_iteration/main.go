package main

import "fmt"

// В силу того, что порядок обхода мапы случайный
// то вывод при работе с этой функцией недетерминированный
// то есть в одном случае будет вывод map[baz:2, foo:0],
// в другом - map[bar:1, baz:2] и так далее.

func main() {
	data := map[string]int{"foo": 0, "bar": 1, "baz": 2}
	for key := range data {
		if key == "foo" {
			delete(data, "bar")
		}
		if key == "bar" {
			delete(data, "foo")
		}
	}

	fmt.Println(data)
}
