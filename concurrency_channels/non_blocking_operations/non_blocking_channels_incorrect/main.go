package main

func tryToReadFromChannel(ch chan string) (string, bool) {
	// len(ch) - возвращает количество элементов в канале.
	// Причем сам код явялется ошибочным, поскольку когда
	// пишется concurrency-код, стоит учитывать, что между
	// двумя любыми строчками кода может встроится другая
	// горутина, которая успеет прочитать из канала.
	if len(ch) != 0 {
		// значение может быть прочитано другой горутиной.
		value := <-ch
		return value, true
	} else {
		return "", false
	}
}

func tryToWriteToChannel(ch chan string, value string) bool {
	// Аналогично, здесь проверяется, что длина меньше capacity,
	// а затем происходит запись в канал. Но при этом между
	// проверкой и записью может встроиться другая горутина.
	// Как следствие заблокируемся при записи.
	if len(ch) < cap(ch) {
		ch <- value
		return true
	} else {
		return false
	}
}
