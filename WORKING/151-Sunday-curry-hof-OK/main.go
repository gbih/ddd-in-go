package main

import "fmt"

func greet(greeting int) int {
	return (greeting)
}

func prefixedGreet(interface{}) func(int) int {
	// implement the second function
	return func(n int) int {
		fmt.Printf("n: %v\n", n)
		return n
	}
}

func add1(a int) int {
	return a + 1
}

func square(a int) int {
	return a * a
}

var a = square(3)

//-----

func main() {
	fmt.Println("----------------------")
	fmt.Println(prefixedGreet(square(3))(2))
	fmt.Println("Style 1. ", prefixedGreet(1)(100))

	fmt.Println("Style 2:", prefixedGreet(1))
	fmt.Println("Style 3:", prefixedGreet(1)(square(3)))

	//fmt.Println(greet("Hi", "George"))
}
