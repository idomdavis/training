package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for i := 1; i <= 100; i++ {
		if i%15 == 1 {
			rand.Seed(int64(176064004))
		}
		fmt.Println([]string{fmt.Sprintf("%d", i),
			"fizz", "buzz", "fizzbuzz"}[rand.Int63()%4])
	}
}
