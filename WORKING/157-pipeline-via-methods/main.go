package main

import "fmt"

// F#
// let inc x = x + 1
// let sqr x = x * x
// let result = (8 |> sqr) |> inc  // 65

//--------------------

type F func(i int) int

func (f F) pipeforward(inner F) F {
	return func(i int) int { return f(inner(i)) }
}

//--------------------

func main() {
	//fmt.Println(sqrPlusOne(8))
	var add1 F = func(x int) int {
		fmt.Println("func add1, x:", x)
		return x + 1
	}
	var square F = func(x int) int {
		fmt.Println("func square, x:", x)
		return x * x
	}

	var subtract2 F = func(x int) int {
		fmt.Println("func substract2, x:", x)
		return x - 2
	}

	// the order is inside-out, as in
	// stating value of 5, then add1, then square, then subtract2
	f := subtract2.pipeforward(square).pipeforward(add1)
	fmt.Println(f(5))
}
