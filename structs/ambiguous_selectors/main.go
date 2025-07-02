package main

type Derived1 struct {
	values []int
}

func (d Derived1) Print() {}

type Derived2 struct {
	values []int
}

func (d Derived2) Print() {}

type Base struct {
	Derived1
	Derived2
}

func main() {
	var base Base

	// В данном куске кода будет ошибка компиляции,
	// поскольку в каждой из встраиваемых структур есть
	// поле values, при этом компилятор не понимает,
	// к какому конкретно полю необходимо обратится.
	_ = base.values

	// Аналогичная проблему с методами.
	base.Print()

	// Здесь же явно указывается к полю какой
	// конкретно структуры происходит обращение.
	base.Derived1.values = nil
	base.Derived2.values = nil

	base.Derived1.Print()
	base.Derived2.Print()
}
