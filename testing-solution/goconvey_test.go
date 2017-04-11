package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test(t *testing.T) {
	Convey("A value divisible by 3", t, func() {
		r := fizzbuzz(6)
		Convey("Will return 'Fizz", func() {
			So(r, ShouldEqual, "Fizzl")

		})
	})
}
