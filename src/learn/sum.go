package learn

func Sum(numbers [5]int) int {
	sum := 0
	for i := 0; i < 5; i++ { //nolint:intrange
		sum += numbers[i]
	}
	return sum
}

func Add(x, y int) int {
	return x + y
}
