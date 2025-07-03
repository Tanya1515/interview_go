package main

import (
	"fmt"
	"os"
)

func ReadFile(filename string) error {
	var err *os.PathError // здесь указатель выставлен в nil
	if filename == "" {
		return err // itab(no nil)/data(nil), поскольку здесь скрывается тип ошибки.
	}

	// reading...
	return err
}

func main() {
	// возвращает ошибку в случае проблем с чтением
	err := ReadFile("")
	if err != nil {
		// При запуске задачи попадем в это условие, поскольку
		fmt.Println("error") // error
	} else {
		fmt.Println("nil")
	}

	fmt.Println("value of err: ", err)        // value of err: <nil>
	fmt.Printf("type of err: %T\n", err)      // type of err: *fs.PathError
	fmt.Println("(err == nil): ", err == nil) // (err == nil): false

}
