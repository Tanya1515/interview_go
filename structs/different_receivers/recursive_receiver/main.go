package main

import "fmt"

// Здесь значение изменится, поскольку мы
// работаем с указателем на вложеннию структуру

type account struct {
	balance int
}

type client struct {
	account *account
}

func (c client) add(value int) {
	c.account.balance += value
}

func main() {
	c := client{
		account: &account{},
	}

	c.add(100)
	fmt.Println(c.account.balance) // 100
}
