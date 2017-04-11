package fizzbuzz_test

import (
	"testing"

	. "github.com/domdavis/training/exercise-8/fizzbuzz"
	. "github.com/smartystreets/goconvey/convey"
)

func Test(t *testing.T) {
	Convey("", t, func() {
		Convey("Will ", func() {
			FizzBuzz()
		})
	})
}
