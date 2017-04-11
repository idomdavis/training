package rn_test

import (
	"fmt"

	"github.com/domdavis/training/rn"
)

func ExampleAtor() {
	fmt.Println(rn.Itor(2017))
	// Output: MMXVII

}

func ExampleRtoa() {
	i, _ := rn.Rtoi("MMXVII")
	fmt.Println(i)
	_, err := rn.Rtoi("IM")
	fmt.Println(err)

	// 2017
	// IM not a valid Roman numeral
}

func ExampleAdd() {
	r, _ := rn.Add("IV", "I")
	fmt.Println(r)
	// Output: V
}

func ExampleSubtract() {
	r, _ := rn.Subtract("X", "I")
	fmt.Println(r)
	// Output: IX
}

func ExampleMultiply() {
	r, _ := rn.Multiply("III", "X")
	fmt.Println(r)
	// Output: XXX
}

func ExampleMod() {
	r, _ := rn.Mod("X", "III")
	fmt.Println(r)
	// Output: I
}

func ExampleDivide() {
	r, _ := rn.Divide("M", "X")
	fmt.Println(r)
	// Output: C
}
