package main

import (
	"fmt"

	"github.com/domdavis/training/exercise-10/fizzbuzz"
)

func main() {
	for i := 0; i < 100; i++ {
		fmt.Printf("%d - %s%s%s\n", i,
			fizzbuzz.FizzBuzz(fizzbuzz.Fizz, i),
			fizzbuzz.FizzBuzz(fizzbuzz.Buzz, i),
			fizzbuzz.FizzBuzz(fizzbuzz.Number, i),
		)
	}
}
