package main

import (
	"fmt"
)

// HOF, function as input
func getAreaFunc() func(int, string, int) int {
	return func(x int, z string, y int) int {
		return x * y
	}
}

// x and y are passed into area(int, int)
func print(x int, z string, y int, area func(int, string, int) int) {
	fmt.Printf("Area is %d\n", area(x, z, y))
}

//---

// HOF function, with another function as input with signature, func(x int) int
func evalWith5ThenAdd2(x int, f func(int) int) int {
	return f(x + 2)
}

func evalWith5ThenAdd20(x int, f func(int) int) int {
	return f(x + 2)
}

// simple function that can be called by itself
func add1(x int) int {
	return x + 1
}

// simple function that can be called by itself
func square(x int) int {
	return x * x
}

//-----

func main() {
	fmt.Println("-----------------------------")
	fmt.Println(add1(0))
	fmt.Println(evalWith5ThenAdd2(2, add1))

	fmt.Println()
	fmt.Println(square(5))
	fmt.Println(evalWith5ThenAdd2(10, square))

}
