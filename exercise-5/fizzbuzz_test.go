package main

import "testing"

func TestFizzBuzz(t *testing.T) {
	for input, want := range map[int]string{
		1:  "1",
		2:  "2",
		3:  "Fizz",
		4:  "4",
		5:  "Buzz",
		15: "FizzBuzz",
	} {
		if r := fizzbuzz(input); r != want {
			t.Error("fizzbuzz(%s), wanted '%s' got %s", input, want, r)
		}
	}
}

func BenchmarkFizzBuzz(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fizzbuzz(i)
	}
}
