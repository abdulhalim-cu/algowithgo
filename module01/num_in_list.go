package module01

// NumInList will return true if number is in the
// list slice, and false otherwise.

func NumInList(list []int, num int) bool {
	for _, val := range list {
		if val == num {
			return true
		}
	}
	return false
}
