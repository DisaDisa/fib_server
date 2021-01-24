package fibcalc

import (
	"errors"
)

func calcFib(x int) int {
	if x <= 1 {
		return x
	}
	ch := make(chan int, 2)
	defer close(ch)
	go func() {
		cache := CacheGet()
		if val, err := cache.GetValue(x - 1); err != nil {
			ch <- val
		} else {
			val := calcFib(x - 1)
			ch <- val
			cache.SetValue(x-1, val)
		}
	}()

	go func() {
		cache := CacheGet()
		if val, err := cache.GetValue(x - 2); err != nil {
			ch <- val
		} else {
			val := calcFib(x - 1)
			ch <- val
			cache.SetValue(x-1, val)
		}
	}()

	return <-ch + <-ch
}

func calcFibSlow(x int) int {
	if x <= 1 {
		return x
	}
	return calcFibSlow(x-1) + calcFibSlow(x-2)
}

//GetFibNimber return n-th number of fibonacci sequence
func GetFibNimber(x int) (int, error) {
	if x < 0 {
		return x, errors.New("Negative value")
	}
	return calcFib(x), nil

}
