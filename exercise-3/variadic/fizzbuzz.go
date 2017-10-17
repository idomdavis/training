package main

import "fmt"

const (
	fizz = "Fizz"
	buzz = "Buzz"
)

func main() {
	fizzbuzz(1, 2, 4, 5, 6, 7, 8, 15, 21, 42)
}

func fizzbuzz(values ...int) {
	for _, i := range values {
		switch {
		case i%3 == 0 && i%5 == 0:
			fmt.Printf("%-5d - %s%s", i, fizz, buzz)
		case i%3 == 0:
			fmt.Printf("%-5d - %s", i, fizz)
		case i%5 == 0:
			fmt.Printf("%-5d - %s", i, buzz)
		default:
			fmt.Print(i)
		}

		fmt.Println()
	}
}
