// https://www.golangprograms.com/higher-order-functions-in-golang.html#:~:text=A%20Higher%2DOrder%20function%20is,arguments%20or%20by%20returning%20them.
package main

import "fmt"

//---------------------------------

//　Returning Functions from other Functions

// Instead of this...
func add1(x int) int { return x + 1 }
func add2(x int) int { return x + 2 }

// Create an "adder generator" or wrapper, a function that
// returns an "add" function with the number to add baked in.

func adderGenerator(numberToAdd int) func(int) int {
	// return a lambda
	return func(x int) int { return numberToAdd + x }
}

func adderGenerator2(numberToAdd int) func(int) int {
	// using a named function instead of anonymous function
	var innerFn = func(x int) int { return numberToAdd + x }
	// return the inner function
	return innerFn
}

// Return multiple functions and one integer value.
// multipleSum func(a int) func(int) func(int) func(int) int
func multipleSum(a int) func(int) func(int) func(int) int {
	return func(b int) func(int) func(int) int {
		return func(c int) func(int) int {
			return func(d int) int {
				return a + b + c + d
			}
		}
	}
}

//---------------------------------

// Passing Functions as Arguments to other Functions
func sum(x, y int) int {
	return x + y
}
func partialSum(x int) func(int) int {
	return func(y int) int {
		return sum(x, y)
	}
}

//---------------------------------

func main() {
	fmt.Println("----------------------------")

	fmt.Println("adderGenerator(1)(2):", adderGenerator(1)(2))
	fmt.Println("adderGenerator2(1):", adderGenerator2(1)(2))

	// With a number "baked in"
	add1 := adderGenerator(1)
	fmt.Println("add1(2):", add1(2))

	add100 := adderGenerator(100)
	fmt.Println("add100(2):", add100(2))

	fmt.Println()

	//　Returning Functions from other Functions
	fmt.Println("multileSum(1)(2)(3)(4):", multipleSum(1)(2)(3)(4))

	add200 := multipleSum(100)
	fmt.Println("add200(1)(1)(1):", add200(1)(1)(1))

	// Passing Functions as Arguments to other Functions
	// partial := partialSum(3)
	// fmt.Println(partial(7))

}
