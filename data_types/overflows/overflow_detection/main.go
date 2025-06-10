package main

import (
	"errors"
	"math"
)

var ErrIntOverflow = errors.New("integer overflow")

// функция, которая позволяет обнаружить переполнение во время
// операции увеличения на 1, аналогично работает декремент
func Inc(counter int) (int, error) {
	if counter == math.MaxInt {
		return 0, ErrIntOverflow
	}

	return counter + 1, nil
}

// функция, которая позволяет обнаружить переполнение при сложении двух чисел
func Add(lhs, rhs int) (int, error) {
	// проверяем знак первого числа (оно положительное)
	if rhs > 0 {
		// сравниваем второе число с разностью максимального числа для этого int-а
		// и первого числа
		if lhs > math.MaxInt-rhs {
			return 0, ErrIntOverflow
		}
		// проверяем, что первое число отрицательное
	} else if rhs < 0 {
		// сравниваем разность минимального числа для этого Int и певрого числа
		// со вторым числом
		if lhs < math.MinInt-rhs {
			return 0, ErrIntOverflow
		}
	}

	return lhs + rhs, nil
}

// функция, которая проверяет переполнение при операции умножения двух чисел
func Mul(lhs, rhs int) (int, error) {

	// проверяем, что хотя бы одно из чисел равно 0
	if lhs == 0 || rhs == 0 {
		return 0, nil
	}

	// проверяем, что хотя бы одно из чисел равно 0
	if lhs == 1 || rhs == 1 {
		return lhs * rhs, nil
	}

	// это условие вызывается по той причине, что модуль от минимального числа
	// для int-переменной на 1 больше, поэтому при умножении на -1 возникнет переполнение
	// другими словами: int8 - 127...-128, то есть при умножении -128*(-1) = 128 - это переполнение
	if (lhs == -1 && rhs == math.MinInt) || (rhs == -1 && lhs == math.MinInt) {
		return 0, ErrIntOverflow
	}

	// делим максимальное/минимальное число для этого Int
	// на одно из чисел и сравниваем со вторым
	if lhs > math.MaxInt/rhs {
		return 0, ErrIntOverflow
	} else if lhs < math.MinInt/rhs {
		return 0, ErrIntOverflow
	}

	return lhs * rhs, nil
}
