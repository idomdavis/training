package main

import (
	"fmt"
	"strconv"
)

const (
	fizz = "Fizz"
	buzz = "Buzz"
)

func main() {
	for _, i := range []int{1, 2, 4, 5, 6, 7, 8, 15, 21, 42} {
		go println(i, fizzbuzz(i))
	}
}

func fizzbuzz(i int) string {
	switch {
	case i%3 == 0 && i%5 == 0:
		return fmt.Sprintf("%s%s", fizz, buzz)
	case i%3 == 0:
		return fizz
	case i%5 == 0:
		return buzz
	default:
		return strconv.Itoa(i)
	}
}
