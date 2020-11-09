package main

import (
	"errors"
	"fmt"
)

// Total Functions
// Restrict the input to eliminate values that are illegal.
// The goal is every input to the function would have a valid output,
// with no exceptions.
// This is the definition of a total function, where every input has
// a corresponding output.

//-----------------------

// https://github.com/golang/go/issues/29649#issuecomment-455081495
// https://play.golang.org/p/JqeeLEIin1_8
// Define Range and Ranged types in Go, and have range checked values
// at run time, with overhead cost of a single pointer per Ranged value.

type TwelveDividedBy func(int) int

type Range struct {
	min int
	max int
}

func NewRange(min, max int) *Range {
	return &Range{min, max}
}

func (r Range) Includes(value int) bool {
	return (value >= r.min) && (value <= r.max)
}
func (r *Range) Check(value int) {
	if r == nil {
		panic("range: Range pointer not valid")
	}
	if !r.Includes(value) {
		panic("range: Value out of range in range check")
	}
}

type Ranged struct {
	value int
	*Range
}

func NewRangedWithRange(value int, ran *Range) Ranged {
	ran.Check(value)
	return Ranged{value, ran}
}

//-----

// constrain to non-zero ints
type NonZeroInteger func(int) (int, error)

var nonZeroInteger NonZeroInteger = func(x int) (int, error) {
	if (x < 0) || (x > 6) {
		fmt.Println("Re-enter a number between 1 and 6")
		return x, errors.New("Number is not between 1 and 6")
	}
	return x, nil
}

// This does NOT implement a Total Function
var twelveDividedBy TwelveDividedBy = func(x int) int {
	switch x {
	case 6:
		return 2
	case 5:
		return 2
	case 4:
		return 3
	case 3:
		return 4
	case 2:
		return 6
	case 1:
		return 12
	case 0:
		panic("Cannot divide by zero")
	default:
		panic("Outside bounds")
	}
}

func main() {
	fmt.Println("------------------------------")
	inputRange := NewRange(1, 6) // &{1 6}

	input2 := NewRangedWithRange(6, inputRange)
	fmt.Println("input2.value:", input2.value)
	fmt.Println(twelveDividedBy(input2.value))

	fmt.Println("===============")

	//--------------

	//fmt.Println(twelveDividedBy(1))
	input, err := nonZeroInteger(6)
	fmt.Println("input:", input)
	if err != nil {
		fmt.Println("Re-enter number")
		return
	}
	fmt.Println(twelveDividedBy(input))

}
