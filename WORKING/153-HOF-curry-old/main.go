// https://blog.usejournal.com/function-currying-in-go-a88672d6ebcf
package main

import "fmt"

// standard two-parameter function
// func(x, y int) int
func add(x, y int) int {
	return x + y
}

// convert into a one-parameter function via currying and
// returning a new function
// func(x int) func(int) int
func adderGenerator(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}

//------------------------

func main() {
	fmt.Println("------------------------------")
	fmt.Println(add(1, 2))
	fmt.Println(adderGenerator(1)(2))

}
