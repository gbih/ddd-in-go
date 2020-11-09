package main

import (
	"fmt"
)

// F#
// let add1 x = x + 1
// letsquare x = x * x

type IntToBool func(x int) bool
type BoolToString func(x bool) string

type Wrapper func(interface{}) interface{}

func Wrapper1(interface{}) interface{} {
	return func(x int) bool {
		return (x % 2) == 0
	}
}

func Wrapper2(interface{}) interface{} {
	return func(x bool) string {
		return fmt.Sprintf("value is %v", x)
	}
}

//----

func main() {
	fmt.Println("-----------------------")

	//fmt.Println(pipeforward(5, isEven, printBool))
	//pipeforward(5, isEven, printBool)

	// Assign functions to array, where we can assign function input
	var arrayFuncs = [...](Wrapper){Wrapper1, Wrapper2}

	var input = 5
	for i, arrayFunc := range arrayFuncs {

		result := (arrayFunc)(input)
		fmt.Printf("\tindex: %v, input: %v, result: %v\n", i, input, result)
		//input = result.(int)
		fmt.Println(input)
		// switch input {
		// case 2:
		// 	fmt.Println(input)
		// }

	}
}
