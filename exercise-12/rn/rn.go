package rn

import (
	"regexp"

	"strings"

	"github.com/domdavis/training/exercise-12/intconv"
)

var RNer intconv.Converter = func(i int) string {
	r := ""
	for j := 0; j < i; j++ {
		r += "I"
	}

	r = regexp.MustCompile(`I{1000}`).ReplaceAllString(r, "M")
	r = regexp.MustCompile(`I{900}`).ReplaceAllString(r, "CM")
	r = regexp.MustCompile(`I{500}`).ReplaceAllString(r, "D")
	r = regexp.MustCompile(`I{400}`).ReplaceAllString(r, "CD")
	r = regexp.MustCompile(`I{100}`).ReplaceAllString(r, "C")
	r = regexp.MustCompile(`I{90}`).ReplaceAllString(r, "XC")
	r = regexp.MustCompile(`I{50}`).ReplaceAllString(r, "L")
	r = regexp.MustCompile(`I{40}`).ReplaceAllString(r, "XL")
	r = regexp.MustCompile(`I{10}`).ReplaceAllString(r, "X")
	r = regexp.MustCompile(`I{9}`).ReplaceAllString(r, "IX")
	r = regexp.MustCompile(`I{5}`).ReplaceAllString(r, "V")
	r = regexp.MustCompile(`I{4}`).ReplaceAllString(r, "IV")
	return r
}

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
