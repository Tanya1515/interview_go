package main

import (
	"errors"
	"fmt"
)

/*
	Оборачивание ошибок - упаковка ошибки в контейнер-обертки, которые
	делают доступной и исходную ошибку.

	Используется при:
	1) Добавлении дополнительного контекста к ошибке.

	2) Маркировке ошибки.

	Для того, чтобы проверить, относится обернутая ошибка к некоторому
	типу или нет - нужно использовать error.As(). Это функция, которая
	рекурсивно разворачивает ошибку.

*/

type DatabaseError struct{}

func (d DatabaseError) Error() string {
	return "database error"
}

func GetDataFromDB() error {
	return fmt.Errorf("failed to get data: %w", DatabaseError{})
}

// Вывод:
// failed to get data: database error

func main() {
	err := GetDataFromDB()
	if errors.As(err, &DatabaseError{}) {
		fmt.Println(err.Error())
	} else {
		fmt.Println("unknown error")
	}
}
