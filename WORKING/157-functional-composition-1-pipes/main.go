package main

import "fmt"

// F#
// let inc x = x + 1
// let sqr x = x * x
// let result = (8 |> sqr) |> inc  // 65

// Go
// pseudo-pipe functionality via applying functions sequentially:
func sqrPlusOne(x int) int {
	sqr := func(x int) int { return x * x }
	inc := func(x int) int { return x + 1 }
	return inc(sqr(x))
}

//--------------------

type F func(i int) int

func (f F) pipeforward(inner F) F {
	return func(i int) int { return f(inner(i)) }
}

//--------------------

func main() {
	//fmt.Println(sqrPlusOne(8))
	var f1 F = func(i int) int { return i * 2 }
	var f2 F = func(i int) int { return i + 1 }
	var f3 F = func(i int) int { return i + 10 }

	f := f1.pipeforward(f2).pipeforward(f3)

	fmt.Println(f(2))
}
