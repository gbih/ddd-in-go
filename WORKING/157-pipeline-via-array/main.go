package main

import (
	"fmt"
)

// F#
// let add1 x = x + 1
// letsquare x = x * x

type Composition func(x int) int

var add1 Composition = func(x int) int {
	fmt.Println("func add1, x:", x)
	return x + 1
}

var square Composition = func(x int) int {
	fmt.Println("func square, x:", x)
	return x * x
}

//----

func main() {
	fmt.Println("-----------------------")

	// Assign functions to array, where we can assign function input
	var arrayFuncs = [...](func(x int) int){add1, square}

	var input = 5
	for i, arrayFunc := range arrayFuncs {
		result := (arrayFunc)(input)
		fmt.Printf("\tindex: %v, input: %v, result: %v\n", i, input, result)
		input = result
	}
}

// func manipulate(action ...Manipulate) Manipulate {
// 	return (fn func(int) int) int {
// 		for _, f := range action {
// 			return fn
// 		}
// 		return y
// 	}
// }
