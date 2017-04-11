package main

import (
	"fmt"

	"github.com/domdavis/training/exercise-9/fizzbuzz"
)

func main() {
	for i := 0; i < 100; i++ {
		fmt.Printf("%d - %s%s%s\n", i,
			fizzbuzz.Fizz.Translate(i),
			fizzbuzz.Buzz.Translate(i),
			fizzbuzz.Number.Translate(i),
		)
	}
}
