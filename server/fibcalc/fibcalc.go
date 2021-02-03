package fibcalc

import "fmt"

//GetFibRange return range from X-th to Y-th of fibonacci sequence
func GetFibRange(x, y int) []int {
	if x > y {
		x, y = y, x
	}
	sequence := make([]int, 0, y-x+1)

	cache := CacheGet()
	if MaxCalculatedIndex < y {
		var prev, cur int
		var err error
		if prev, err = cache.GetValue(MaxCalculatedIndex - 1); err != nil {
			return getFibRangeSlow(x, y)
		}
		if cur, err = cache.GetValue(MaxCalculatedIndex); err != nil {
			return getFibRangeSlow(x, y)
		}
		for i := MaxCalculatedIndex + 1; i <= y; i++ {
			fmt.Println("FibCalc", prev, cur)
			newVal := cur + prev
			prev = cur
			cur = newVal

			if err = cache.SetValue(i, cur); err != nil {
				return getFibRangeSlow(x, y)
			}
		}
	}
	for i := x; i <= y; i++ {
		var val int
		var err error
		if val, err = cache.GetValue(i); err != nil {
			return getFibRangeSlow(x, y)
		}
		sequence = append(sequence, val)
	}
	return sequence
}

func getFibRangeSlow(x, y int) []int {
	if x > y {
		x, y = y, x
	}
	sequence := make([]int, 0, y-x+1)
	prev := 0
	cur := 1
	if x == 1 {
		sequence = append(sequence, 0)
		if y >= 2 {
			sequence = append(sequence, 1)
		}
	}
	if x == 2 {
		sequence = append(sequence, 1)
	}
	for i := 3; i <= y; i++ {
		newVal := cur + prev
		prev = cur
		cur = newVal
		if i >= x {
			sequence = append(sequence, cur)
		}
	}
	return sequence
}
