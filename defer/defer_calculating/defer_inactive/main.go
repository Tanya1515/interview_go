package main

import "fmt"

/*

 Стоит отметить, что если в defer располагается код,
 который недоступен для выполнения (unreachable),
 то функция просто не будет вызвана.


 То есть здесь будет напечатано слово code.
*/

func main() {
	defer fmt.Println("code")
	if false {
		defer fmt.Println("unreacheable code")
	}

	return
	defer fmt.Println("unreacheable code")
}
