package main

import (
	"errors"
	"fmt"
	"io"
)

// Однако, в Golang существуют специализированные
// константные ошибки, которые нельзя менять.

// Здесь создаем некоторый type definition для строки
type Error string

// Для этого type definition создается специальный метод Error,
// который возвращает строку.

func (e Error) Error() string {
	// В силу того, что базовый тип - строка,
	// каст к строке успешно проходит.
	return string(e)
}

// Создаем константную строку по сути
const ErrEOF = Error("EOF")

func main() {
	var err error = ErrEOF // типу error присваивается константная строка ErrEOF.
	// затем в err записывается другая ошибка
	// и эта операция производится корректно,
	// поскольку err типа error, а это - интерфейсный тип.
	err = io.EOF
	_ = err

	// ErrEOF = Error("new error") -> coplilation error, поскольку ErrEOF - константа

	anotherErrEOF1 := Error("EOF")
	// эти две ошибки будут равны, поскольку константные ошибки
	// по своей сути - это константные строки.
	fmt.Println(`anotherErrEOF1 = Error("EOF"):`, anotherErrEOF1 == ErrEOF) // true

	// несмотря на то, что тексты ошибок равны,
	// сами ошибки равны не будут, поскольку это абсолютно
	// разные ошибки и объекты сами по себе
	anotherErrEOF2 := errors.New("EOF")
	fmt.Println("anotherErrEOF2 = io.EOF:", anotherErrEOF2 == io.EOF) // false
}
