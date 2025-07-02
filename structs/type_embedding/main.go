package main

import "fmt"

type Person struct {
	Name string
}

func (p *Person) Intro() string {
	return p.Name
}

// Встраивание - концепция, связанная с перекрытием методов.
// Встраивание (embedding) — это особый вид композиции в Go,
// при котором тип (обычно структура) включается в другую
// структуру без явного имени поля. Это позволяет автоматически
// "поднимать" (promote) методы и поля встроенного типа в область
// видимости внешней структуры. То есть если бы не было метода Intro(),
// woman.Intro() все равно бы выполнилось, поскольку был бы взят метод
// у типа Person.

type Woman struct {
	Person
}

func (w *Woman) Intro() string {
	return "Mrs. " + w.Person.Intro()
}

// Композиция в Go — это общий принцип проектирования, при котором
// структуры (structs) включают другие структуры или типы как поля
// для использования их функциональности. Доступ к методам и полям
// вложенной структуры осуществляется через имя поля. Композиция
// не подразумевает автоматического делегирования методов.
// То есть вызов Man.Intro работать не будет.

type Man struct {
	person Person
}

func main() {
	woman := Woman{
		Person: Person{
			Name: "Ekaterina",
		},
	}

	fmt.Println(woman.Intro()) // - Mrs. Ekaterina
	//fmt.Println(woman.Person.Intro()) - Ekaterina
}
