package exercise_9_test

import (
	"fmt"

	rn "github.com/domdavis/training/exercise-9"

	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func ExampleToRomanNumeral() {
	fmt.Println(rn.ToRomanNumeral(2017))

	// Output:
	// MMXVII
}

func ExampleToRomanNumeral2() {
	fmt.Println(rn.ToRomanNumeral(10000))

	// Output:
	// MMMMMMMMMM
}

func ExampleToInt() {
	fmt.Println(rn.ToInt("MMXVII"))

	// Output:
	// 2017
}

func ExampleToInt2() {
	fmt.Println(rn.ToInt("MMXVMII"))

	// Output:
	// 2015
}

func TestToRomanNumeral(t *testing.T) {
	Convey("A roman numeral", t, func() {
		n := "MMXVII"
		i := rn.ToInt(n)

		Convey("Will convert to 2017", func() {
			So(i, ShouldEqual, 2017)
		})
	})
}
