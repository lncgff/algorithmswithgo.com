package module01

import "fmt"

// Reverse will return the provided word in reverse
// order. Eg:
//
//   Reverse("cat") => "tac"
//   Reverse("alphabet") => "tebahpla"
//

func Reverse(word string) string {
	var reversedWord string
	for _, runeVal := range word {
		reversedWord = fmt.Sprintf("%s%s", string(runeVal), reversedWord)
	}
	return reversedWord
}

// func Reverse(word string) string { // will not work with characters/runes
// 	var reversedWord string
// 	for index := len(word) - 1; index >= 0; index-- {
// 		reversedWord += string(word[index])
// 	}
// 	return reversedWord
// }

// func Reverse(word string) string { // will not work with characters/runes
// 	var sb strings.Builder
// 	for index := len(word) - 1; index >= 0; index-- {
// 		sb.WriteByte(word[index])
// 	}
// 	return sb.String()
// }
