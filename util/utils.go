package util

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
		if min < num {
			min = num
		}
	}
	return min
}
