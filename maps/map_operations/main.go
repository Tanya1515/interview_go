package main

// При чтении из nil-map-ы код отработает корректно.
func readFromNilMap() {
	var data map[int]int
	_ = data[100]
}

// При удалении ключа из nil-map-ы не будет никаких ошибок.
func deleteFromNilMap() {
	var data map[int]int
	delete(data, 100)
}

// При записи ключа в nil-map-у будет ошибка в рантайме,
// поскольку мапа не проинициализирована.
func writeToNilMap() {
	var data map[int]int
	data[100] = 100
}

// При прохождении range-ом по nil-map-е ничего не произойдет.
func rangeByNilMap() {
	var data map[int]int
	for range data {
	}
}

// При попытке перезаписать существующее значение -
// код также будет работать корректно.
func rewriteExistingKey() {
	data := make(map[int]int)
	data[100] = 500
	data[100] = 1000
}

// При удалении несуществующего ключа из nil-map-ы код
// также отработает корректно.
func deleteNonExistingKey() {
	data := make(map[int]int)
	delete(data, 100)
}

func main() {
	deleteNonExistingKey()
}
