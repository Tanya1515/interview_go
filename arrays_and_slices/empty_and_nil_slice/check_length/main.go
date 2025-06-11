package main

// Need to show solution

// В силу разницы между пустым и нулевым срезом
// возникает проблема, поскольку сравниваться с nil-ом
// не всегда корректно. По этой причине лучше проверять длину среза
// она и для пустого и для nil-ого среза будлет равна 0.

func handleOperations(id string) {
	operations := getOperations(id)
	if operations == nil { // здесь проверка будет некорректно работать для пустого среза
		// handling...
	}
}

func getOperations(id string) []float32 {
	opearations := []float32{} // здесь создается пустой срез, но он не nil.
	if id == "" {
		return opearations
	}

	// adding operations...
	return opearations
}
