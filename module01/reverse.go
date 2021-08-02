package module01

import "strings"

// Reverse takes a string and return its reverse.
// Approach 1: converting string into runes then swapping their first and last position
// func Reverse(word string) string {
// 	runes := []rune(word)
// 	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
// 		runes[i], runes[j] = runes[j], runes[i]
// 	}
// 	return string(runes)
// }

// Approach 2: declaring an emtpy stirng then add char to its first position
// func Reverse(word string) string {
// 	newString := ""
// 	for _, char := range word {
// 		newString = string(char) + newString
// 	}
// 	return newString
// }

// Approach 3: build new string from backorder
// func Reverse(word string) string {
// 	var newString string
// 	for i := len(word) - 1; i >= 0; i-- {
// 		newString = newString + string(word[i])
// 	}
// 	return newString
// }

// Approach 4: Using string builder
func Reverse(word string) string {
	var stringBuilder strings.Builder
	for i := len(word) - 1; i >= 0; i-- {
		stringBuilder.WriteByte(word[i])
	}
	return stringBuilder.String()
}
