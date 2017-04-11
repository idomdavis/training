package rn

import (
	"regexp"

	"github.com/domdavis/training/exercise-10/intconv"
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
