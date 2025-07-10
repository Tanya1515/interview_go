package main

import (
	"errors"
	"fmt"
)

// Для того, чтобы проверить, относится ли значение обернутой
// ошибки к определенному значению - нужно использовать errors.Is().
// Эта функция рекурсивно разворачивает ошибку и проверяет ошибку.

var ErrDatabaseProblem = errors.New("database problem")

func GetDataFromDB() error {
	return fmt.Errorf("failed to get data: %w", ErrDatabaseProblem)
}

func main() {
	err := GetDataFromDB()
	if errors.Is(err, ErrDatabaseProblem) {
		fmt.Println(err.Error())
	} else {
		fmt.Println("unknown error")
	}
}
