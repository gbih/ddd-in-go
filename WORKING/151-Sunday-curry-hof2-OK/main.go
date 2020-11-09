package main

import "fmt"

func greet(greeting int) int {
	return (greeting)
}

func prefixedGreet(closure int) func(int) int {
	// implement the second function
	return func(n int) int {
		fmt.Printf("closure: %v, n: %v\n", closure, n)
		return n + closure
	}
}

func add1(a int) int {
	return a + 1
}

func square(a int) int {
	return a * a
}

//-----

func main() {
	//fmt.Println("Style 1. ", prefixedGreet(1)(100))
	fmt.Println("----------------------")
	fmt.Println("Style 2:", prefixedGreet(1)(add1(3)))
	fmt.Println("Style 3:", prefixedGreet(1)(square(3)))

	//fmt.Println(greet("Hi", "George"))
}
