package module01

// DecToBase will return a string representing
// the provided decimal number in the provided base.
// This is limited to bases 2-16 for simplicity.
//
// Eg:
//
//   DecToBase(14, 16) => "E"
//   DecToBase(14, 2) => "1110"
//
// remainder --> append --> divide --> remainder --> append etc...
func DecToBase(dec, base int) string {
	const charset = "0123456789ABCDEF"
	var converted string
	for dec > 0 {
		rem := dec % base
		converted = string(charset[rem]) + converted
		dec = dec / base
	}
	return converted
}

// func remainder(dec, base int) string {
// 	rem := dec % base
// 	switch rem {
// 	case 10:
// 		return "A"
// 	case 11:
// 		return "B"
// 	case 12:
// 		return "C"
// 	case 13:
// 		return "D"
// 	case 14:
// 		return "E"
// 	case 15:
// 		return "F"
// 	default:
// 		return strconv.Itoa(rem)
// 	}
// }
