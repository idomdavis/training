package main

import (
	"fmt"
	"math/rand"
)

func main() {
	max := 0
	expect := []int{0, 0, 1, 0, 2, 1, 0, 0, 1, 2, 0, 1, 0, 0, 3}
	seed := int64(0)
	wanted := int64(0)

	for wanted == int64(0) {
		rand.Seed(seed)
		i := 0
		good := true

		for i < 15 && good {
			r := rand.Intn(3)
			if expect[i] != r {
				good = false
			} else if i > max {
				max = i
				fmt.Printf("==> %d %v\n", max, expect[:max])
			}
			i++
		}

		seed++
		if good {
			wanted = seed
		}
	}

	fmt.Println(wanted)

}
