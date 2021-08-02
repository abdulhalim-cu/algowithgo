package module01

// Sum takes a slice of integers and will return sum of all
// numbers in the list or 0 for empty list

func Sum(numbers []int) int {
	sum := 0
	for _, val := range numbers {
		sum += val
	}
	return sum
}
