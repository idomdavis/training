package main

import (
	"fmt"
	"strconv"
)

type converter func(in <-chan fizzbuzz, out chan<- fizzbuzz)

type fizzbuzz struct {
	in  int
	out string
}

var converters = []converter{fizz, buzz, number}

func main() {
	chans := make([]chan fizzbuzz, 4)

	for i := range chans {
		chans[i] = make(chan fizzbuzz)
		defer close(chans[i])
	}

	for i := 0; i < len(chans)-1; i++ {
		go converters[i](chans[i], chans[i+1])
	}

	for i := 1; i < 100; i++ {
		chans[0] <- fizzbuzz{in: i}
		r := <-chans[len(chans)-1]
		fmt.Printf("%d - %s\n", r.in, r.out)
	}
}

func fizz(in <-chan fizzbuzz, out chan<- fizzbuzz) {
	for i := range in {
		if i.in%3 == 0 {
			i.out = "Fizz"
		}
		out <- i
	}
}

func buzz(in <-chan fizzbuzz, out chan<- fizzbuzz) {
	for i := range in {
		if i.in%5 == 0 {
			i.out += "Buzz"
		}
		out <- i
	}
}

func number(in <-chan fizzbuzz, out chan<- fizzbuzz) {
	for i := range in {
		if i.in%3 != 0 && i.in%5 != 0 {
			i.out = strconv.Itoa(i.in)
		}
		out <- i
	}
}


