package main

import "fmt"

// F#
// let isEven x = (x % 2) = 0 // an int -> bool function
// let printBool x = sprintf "value is %b" x // a bool -> string function
// let isEvenThenPrint x = x |> isEven |> printBool // int -> string function
// isEvenThenPrint 2 // result is "value is true"

// Go
// pseudo-pipe functionality via applying compatible functions sequentially:

// func(x int) bool
func isEven(x int) bool { return (x % 2) == 0 }

// func(x bool) string
func printBool(x bool) string { return fmt.Sprintf("value is %v", x) }

// func(x int) string
func isEvenThenPrint(x int) string {
	//isEven := func(x int) bool { return (x % 2) == 0 }
	//printBool := func(x bool) string { return fmt.Sprintf("value is %v", x) }
	return printBool(isEven(x))
}

func main() {
	fmt.Println(isEvenThenPrint(2))
}
