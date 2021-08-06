package module01

import "fmt"

func DecToBase(dec, base int) string {
	var newNumber string
	for dec > 0 {
		remainder := dec % base
		newNumber = fmt.Sprintf("%X%s", remainder, newNumber)
		dec = dec / base
	}
	return newNumber
}
