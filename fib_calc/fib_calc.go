package fib_calc

import "errors"

func calcFib(x int) int {
	if x <= 1 {
		return x
	}
	return calcFib(x-1) + calcFib(x-2)
}

func GetFibNimber(x int) (int, error) {
	if x < 0 {
		return x, errors.New("Negative value")
	}
	return calcFib(x), nil
}
