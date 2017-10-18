package fizzbuzz

import (
	"strconv"

	"github.com/domdavis/training/exercise-12/intconv"
)

var FizzBuzzer intconv.Converter = func(i int) string {
	switch {
	case i%3 == 0 && i%5 == 0:
		return "FizzBuzz"
	case i%3 == 0:
		return "Fizz"
	case i%5 == 0:
		return "Buzz"
	default:
		return strconv.Itoa(i)
	}
}
