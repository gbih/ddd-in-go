package main

import (
	"fmt"
)

// F#
// let plus3 x = x + 3 // plus3 : x:int -> int
// let times2 x = x * 2 // times2 : x:int -> int
// letsquare=(funx->x*x) // square : x:int->int
// let addThree = plus3 // addThree : (int -> int)

// Go using function signatures for more constraint
type Plus3 func(x int) int

var plus3 Plus3 = func(x int) int {
	return x + 3
}

type Times2 func(x int) int

var times2 Times2 = func(x int) int {
	return x * 2
}

type Square func(x int) int

var square Square = func(x int) int {
	return x * x
}

type AddThree Plus3

// Constrain function signature to Plus3, but can implement freely
var addThree AddThree = func(x int) int {
	return x + 3
}

func George(x int) int {
	return x * 10000
}

//----

func main() {
	fmt.Println("-----------------------")
	//fmt.Printf("%v\n", addThree(3))
	///

	// Assign functions to array
	var arrayFuncs = [...](func(x int) int){George, addThree, times2, square}
	//fmt.Println(arrayFuncs)

	for _, arrayFunc := range arrayFuncs {
		result := (arrayFunc)(100)
		fmt.Println(result)
	}
}
