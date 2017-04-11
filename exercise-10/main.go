package main

import (
	"fmt"

	"github.com/domdavis/training/exercise-10/fizzbuzz"
	"github.com/domdavis/training/exercise-10/intconv"
	"github.com/domdavis/training/exercise-10/rn"
)

func main() {
	fmt.Println(fizzbuzz.FizzBuzzer.Convert(1))
	fmt.Println(rn.RNer.Convert(2))

	fmt.Println(intconv.Convert(fizzbuzz.FizzBuzzer, 3))
	fmt.Println(intconv.Convert(fizzbuzz.FizzBuzzer, 4))

	fmt.Println(fizzbuzz.FizzBuzzer(5))
	fmt.Println(rn.RNer(6))
}
