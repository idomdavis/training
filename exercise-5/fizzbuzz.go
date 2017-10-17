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
	for k, v := range fizzbuzz(1, 2, 4, 5, 6, 7, 8, 15, 21, 42) {
		println(k, v)
	}
}

func fizzbuzz(values ...int) map[int]string {
	m := make(map[int]string, len(values))

	for _, i := range values {
		switch {
		case i%3 == 0 && i%5 == 0:
			m[i] = fmt.Sprintf("%s%s", fizz, buzz)
		case i%3 == 0:
			m[i] = fizz
		case i%5 == 0:
			m[i] = buzz
		default:
			m[i] = strconv.Itoa(i)
		}
	}

	return m
}
