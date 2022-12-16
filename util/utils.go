package util

func Clamp(min, val, max int) int {
	return Max(min, Min(val, max))
}

func Max(a int, rest ...int) int {
	if len(rest) == 0 {
		return a
	}
	max := a
	for _, num := range rest {
		if num > max {
			max = num
		}
	}
	return max
}

func Min(a int, rest ...int) int {
	if len(rest) == 0 {
		return a
	}
	min := a
	for _, num := range rest {
		if num < min {
			min = num
		}
	}
	return min
}
