package main

import (
	"fmt"
)

type TwelveDividedBy func(int) int

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
	fmt.Println(twelveDividedBy(1))
}
