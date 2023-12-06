package lib

func IntMax(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func IntSum(ints []int) int {
	sum := 0
	for _, i := range ints {
		sum += i
	}
	return sum
}
