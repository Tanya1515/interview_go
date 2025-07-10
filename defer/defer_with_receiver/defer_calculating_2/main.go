package main

type Data struct{}

func MakeData(pointer *int) Data {
	println("MakeData:", *pointer)
	return Data{}
}

func (Data) Print(pointer *int) {
	println("Print:", *pointer)
}

func (d Data) TestPrint(pointer *int) Data {
	println("Print:", *pointer)
	return d
}

// Вывод функции:
// MakeData: 1
// MakeData: 0
// Print: 2

func main() {
	var value = 1
	var pointer = &value
	// Когда отклываем выполнение функции Print,
	// к этой функции необходимо подобраться,
	// поэтому выполняется конструктор Data - MakeData,
	// в процессе которого печатается MakeData: 1
	// (первое значение pointer-а).
	defer MakeData(pointer).Print(pointer)

	/* Если же отложить вызов вот такой функции:
	defer MakeData(pointer).TestPrint(pointer).Print(pointer)
	То в отложенный вызов необходимо добавить reciver, который нужно получить,
	Таким образом, будет выполнена функция TestPrint со значением 1. То есть
	MakeData и TestPrint не откладываются, а откладывается только последняя функция.
	*/

	// Value меняется на 2. Затем
	// переприсваивается значение pointer-а,
	// значение которого теперь 0. Но в момент, когда
	// откладывалась функция Print, туда был скопирован
	// указатель на value, поэтому будет напечатана 2.
	value = 2
	pointer = new(int)
	MakeData(pointer)
}
