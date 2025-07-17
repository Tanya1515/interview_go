package main

import "sync"

// Есть некоторая структура данных с mutex и срезом int.
type Data struct {
	sync.Mutex
	values []int
}

// Функция Add блочит mutex
func (d *Data) Add(value int) {
	d.Lock()
	defer d.Unlock()

	d.values = append(d.values, value)
}

func main() {
	data := Data{}
	data.Add(100)

	// Поскольку в структуре Data используется встраивание
	// типов ничего не мешает просто вызвать Lock/UnLock/TryLock
	// и так далее. В силу того что здесь встроен тип mutex,
	// то здесь нарушена инкапсуляция. По этой причине лучше
	// инкапсулировать мьютекс, то есть ввести дополнительное поле mtx.

	data.Unlock() // Possible problem!
}
