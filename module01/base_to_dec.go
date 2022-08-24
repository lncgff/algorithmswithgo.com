package module01

import "strings"

// BaseToDec takes in a number and the base it is currently
// in and returns the decimal equivalent as an integer.
//
// Eg:
//
//   BaseToDec("E", 16) => 14
//   BaseToDec("1110", 2) => 14
//
func BaseToDec(value string, base int) int {
	const charset = "0123456789ABCDEF"
	var result int
	multiplier := 1
	// loop through value string backwards
	for i := len(value) - 1; i >= 0; i-- {
		pos := strings.Index(charset, string(value[i]))
		result += pos * multiplier
		multiplier *= base
	}
	return result
}
