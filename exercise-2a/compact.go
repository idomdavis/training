package main

import "fmt"

var pattern = []int{0, 0, 1, 0, 2, 1, 0, 0, 1, 2, 0, 1, 0, 0, 3}

func main() {
	for i := 1; i <= 100; i++ {
		fmt.Println([]string{fmt.Sprintf("%d", i),
			"fizz", "buzz", "fizzbuzz"}[pattern[(i-1)%15]])
	}
}
