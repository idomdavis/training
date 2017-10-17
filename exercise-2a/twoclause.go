package main

import "fmt"

const (
	fizz = "Fizz"
	buzz = "Buzz"
)

func main() {
	for i := 1; i <= 100; i++ {
		txt := ""

		if i%3 == 0 {
			txt += fizz
		}

		if i%5 == 0 {
			txt += buzz
		}

		fmt.Printf("%-5d - %s\n", i, txt)
	}
}
