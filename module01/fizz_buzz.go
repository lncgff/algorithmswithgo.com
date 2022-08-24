package module01

import (
	"fmt"
	"strings"
)

// FizzBuzz will print out all of the numbers
// from 1 to N replacing any divisible by 3
// with "Fizz", and divisible by 5 with "Buzz",
// and any divisible by both with "Fizz Buzz".
//
// Note: The test for this is a little
// complicated so that you can just use the
// `fmt` package and print to standard out.
// I wouldn't normally recommend this, but did
// it here to make life easier for beginners.
func FizzBuzz(n int) {
	var elements []string
	for num := 1; num <= n; num++ {
		switch {
		case (num%3 == 0 && num%5 == 0):
			elements = append(elements, "Fizz Buzz")
		case num%3 == 0:
			elements = append(elements, "Fizz")
		case num%5 == 0:
			elements = append(elements, "Buzz")
		default:
			elements = append(elements, fmt.Sprintf("%d", num))
		}
	}
	fmt.Println(strings.Join(elements, ", "))
}
