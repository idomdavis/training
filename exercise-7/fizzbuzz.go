package main

import (
	"fmt"
	"strconv"
)

const (
	Fizz = "Fizz"
	Buzz = "Buzz"
)

var fizz func(int) string = func(i int) string {
	if i%3 == 0 {
		return Fizz
	}

	return ""
}

var buzz func(int) string = func(i int) string {
	if i%5 == 0 {
		return Buzz
	}

	return ""
}

var number func(int) string = func(i int) string {
	if i%3 != 0 && i%5 != 0 {
		return strconv.Itoa(i)
	}

	return ""
}

func main() {
	for _, i := range []int{1, 2, 4, 5, 6, 7, 8, 15, 21, 42} {
		fmt.Printf("%d - %s%s%s\n", i, fizz(i), buzz(i), number(i))
	}
}
