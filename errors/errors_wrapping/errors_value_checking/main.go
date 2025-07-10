package main

import (
	"errors"
	"fmt"
)

// Иногда нужно проверять не тип ошибки, а их значение

var ErrDatabaseProblem = errors.New("database problem")

func GetDataFromDB() error {
	return fmt.Errorf("failed to get data: %w", ErrDatabaseProblem)
}

// Вывод: unknown error

func main() {
	err := GetDataFromDB()
	if err == ErrDatabaseProblem {
		fmt.Println(err.Error())
	} else if err != nil {
		fmt.Println("unknown error")
	}
}
