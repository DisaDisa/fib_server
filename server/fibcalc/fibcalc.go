package fibcalc

//GetFibNimber return range from X-th to Y-th of fibonacci sequence
func GetFibRange(x, y int) []int {
	if x > y {
		x, y = y, x
	}
	sequence := make([]int, y-x+1)
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

	cache := CacheGet()
	for i := 3; i <= y; i++ {
		if val, err := cache.GetValue(i); err == nil {
			sequence = append(sequence, val)
		} else {
			newVal := cur + prev
			prev = cur
			cur = newVal
			sequence = append(sequence, cur)
			go cache.SetValue(i, cur)
		}
	}
	return sequence
}
