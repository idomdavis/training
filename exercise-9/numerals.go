// This package handles conversion to and from Roman Numerals. It uses the
// Wikipedia definition for Roman numerals and therefore may not handle as
// complete a set as other implementations.
package exercise_9

import (
	"fmt"
	"regexp"
	"strings"
)

var valid = regexp.MustCompile(`M?(CM)?D?(CD)?(C)?(XC)?(L)?(XL)?(X)?(IX)?(V)?(IV)?(I)?`)

// Itor converts the provided integer to a Roman numeral string. The provided
// integer must be greater than 0.
func Itor(v int) string {
	r := ""
	for i := 0; i < v; i++ {
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

// Rtoi converts the provided Roman numeral string to an integer. If the string
// is not a valid set of Roman numerals it returns an error.
func Rtoi(r string) (int, error) {

	if !valid.MatchString(r) {
		return 0, fmt.Errorf("%s is not a valid Roman numeral", r)
	}

	i := strings.Count(r, "CM") * 900
	r = strings.Replace(r, "CM", "", -1)
	i += strings.Count(r, "M") * 1000
	r = strings.Replace(r, "M", "", -1)
	i += strings.Count(r, "CD") * 400
	r = strings.Replace(r, "CD", "", -1)
	i += strings.Count(r, "D") * 500
	r = strings.Replace(r, "D", "", -1)
	i += strings.Count(r, "XC") * 90
	r = strings.Replace(r, "XC", "", -1)
	i += strings.Count(r, "C") * 100
	r = strings.Replace(r, "C", "", -1)
	i += strings.Count(r, "XL") * 40
	r = strings.Replace(r, "XL", "", -1)
	i += strings.Count(r, "L") * 50
	r = strings.Replace(r, "L", "", -1)
	i += strings.Count(r, "IX") * 9
	r = strings.Replace(r, "IX", "", -1)
	i += strings.Count(r, "X") * 10
	r = strings.Replace(r, "X", "", -1)
	i += strings.Count(r, "IV") * 4
	r = strings.Replace(r, "IV", "", -1)
	i += strings.Count(r, "V") * 5
	r = strings.Replace(r, "V", "", -1)
	i += strings.Count(r, "I") * 1
	r = strings.Replace(r, "I", "", -1)

	return i, nil
}

func Add(a, b string) (string, error) {
	return calc(a, b, func(a, b int) int { return a + b })
}

func Subtract(a, b string) (string, error) {
	return calc(a, b, func(a, b int) int { return a - b })
}

func Multiply(a, b string) (string, error) {
	return calc(a, b, func(a, b int) int { return a * b })
}

func Mod(a, b string) (string, error) {
	return calc(a, b, func(a, b int) int { return a % b })
}

func Divide(a, b string) (string, error) {
	return calc(a, b, func(a, b int) int { return a / b })
}

func calc(a, b string, f func(int, int) int) (string, error) {
	res := ""
	l, err := Rtoi(a)

	if err != nil {
		return res, err
	}

	r, err := Rtoi(b)

	if err != nil {
		return res, err
	}

	return Itor(f(l, r)), nil
}
