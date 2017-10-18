package exercise_9

import "strings"

var values = []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
var numerals = []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IV",
	"V", "IV", "I"}

// Function ToInt converts a well formatted Roman Numeral string into an
// integer. If the input string is not well formatted then it will convert as
// much of the string as possible, but will give up on the first invalid
// character. No errors or warnings are output when this occurs
func ToInt(in string) int {
	t := 0
	for i, n := range numerals {
		for strings.HasPrefix(in, n) {
			t += values[i]
			in = strings.TrimPrefix(in, n)
		}
	}

	return t
}

// Function ToRomanNumeral converts an integer into its equivalent Roman
// Numeral. The largest recognised numeral is "M" for 1000, so very large
// numbers will generate a large number of "M"s in the output.
func ToRomanNumeral(in int) string {
	var numeral, result string

	for in > 0 {
		numeral, in = convert(in)
		result += numeral
	}

	return result
}

func convert(in int) (string, int) {
	for i, v := range values {
		if in >= v {
			return numerals[i], in - v
		}
	}

	return "", 0
}
