package main

import "fmt"

// F#
// let evalWith5ThenAdd2 fn = fn 5 + 2;;
// val evalWith5ThenAdd2 : fn : (int -> int) -> int

// Go
// fn function takes a function as a argument
func evalWith5ThenAdd2(fn func(int) int) int {
	return fn(5) + 2
}

func add1(x int) int {
	fmt.Println("add1 x:", x)
	return x + 1
}

func square(x int) int {
	fmt.Println("square x:", x)
	return x * x
}

//---------

// TO-DO
//let num = 10
//let squareIt = fun n -> n * n
//System.Console.WriteLine(squareIt num)
//func Sphere(vol func(radius float64) float64) {
// func squareIt(vol func(radius int) int) interface{} {
// 	var a int
// 	return vol(a)
// }

// func num(radius int) int {
// 	return radius
// }

//-----------------------------

// Main Function
func main() {
	fmt.Println("---------------")
	fmt.Println(evalWith5ThenAdd2(add1))
	fmt.Println(evalWith5ThenAdd2(square))
}

/*
HOF
https://www.geeksforgeeks.org/higher-order-function-in-golang/
https://swlaschin.gitbooks.io/fsharpforfunandprofit/content/posts/how-types-work-with-functions.html
https://docs.microsoft.com/en-us/dotnet/fsharp/introduction-to-functional-programming/first-class-functions
There are 2 types:
1. Function as input
2. Function as outpu
*/
