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

func IntProduct(ints []int) int {
	product := 1
	for _, i := range ints {
		product *= i
	}
	return product
}

func IsNeighbour(x0 int, y0 int, x1 int, y1 int) bool {
	return x1 >= x0-1 && x1 <= x0+1 && y1 >= y0-1 && y1 <= y0+1
}

type IntSet struct {
	values []int
}

func (is *IntSet) contains(value int) bool {
	for _, v := range is.values {
		if v == value {
			return true
		}
	}
	return false
}

func (is *IntSet) Add(value int) {
	if !is.contains(value) {
		is.values = append(is.values, value)
	}
}

func (is *IntSet) Sum() int {
	return IntSum(is.values)
}
