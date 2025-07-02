package main

import "fmt"

type Data struct{}

func (d Data) Print() {
	fmt.Println("data")
}

func main() {
	var data Data

	data.Print()    // data
	(&data).Print() // data

	(Data).Print(data)   // data
	(*Data).Print(&data) // data
}
