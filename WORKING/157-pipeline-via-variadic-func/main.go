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

var subtract2 Composition = func(x int) int {
	fmt.Println("func substract2, x:", x)
	return x - 2
}

//----

func pipeforward(input int, action ...Composition) int {
	result := 0
	for _, arrayFunc := range action {
		result = arrayFunc(input)
		fmt.Printf("\tinput: %v, result: %v\n", input, result)
		input = result
	}
	return result
}

func main() {
	fmt.Println("-----------------------")

	fmt.Println(pipeforward(5, add1, square, subtract2))

}
