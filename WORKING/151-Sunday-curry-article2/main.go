// https://blog.usejournal.com/function-currying-in-go-a88672d6ebcf
package main

import "fmt"

type Multiply func(int) func(int) int

var multiply Multiply = func(a int) func(int) int {
	//func multiply(a int) func(int) int {
	return func(i int) int {
		return a * i
	}
}

func subtract(a int) func(int) int {
	return func(i int) int {
		return i - a
	}
}

func main() {

	in := 100
	multiply4 := multiply(4)
	subtract10 := subtract(10)

	// subtract, then multiply
	// Go applies the inner function first.

	fmt.Println(func(i int) int { return multiply4(subtract10(i)) }(100))
	fmt.Println(func(i int) int { return subtract10(multiply4(i)) }(100))

	fmt.Println(func(i int) int {
		// explicitly working out this function:
		// return multiply(4)(subtract(10)(i))

		inner := (subtract(10)(i))
		outer := multiply(4)(inner)
		fmt.Println("i:", i)
		fmt.Println("inner: ", inner)
		fmt.Println("outer: ", outer)
		return outer

	}(100))

	sm := func(i int) int { return multiply4(subtract10(i)) }
	ms := func(i int) int { return subtract10(multiply4(i)) }
	fmt.Printf("%v %v\n", ms(in), sm(in))
}

/*
f# https://fsharpforfunandprofit.com/posts/function-signatures/
Defining your own types for function signatures
Sometimes you may want to create your own types to match a desired
function signature. You can do this using the “type” keyword, and
define the type in the same way that a signature is written.
You can then use these types to constrain function values and
parameters.
*/
