// https://stackoverflow.com/questions/19394868/how-can-go-lang-curry
// pseudo-currying in Go
package main

import "fmt"

func mkAdd(a int) func(int) int {
	return func(b int) int {
		return a + b
	}
}

//-----

type Add func(int, int) int

func (f Add) Apply(i int) func(int) int {
	return func(j int) int {
		return f(i, j)
	}
}

//-----

// https://blog.usejournal.com/function-currying-in-go-a88672d6ebcf

func multiply(a int) func(int) int {
	return func(i int) int {
		return a * i
	}
}
func subtract(a int) func(int) int {
	return func(i int) int {
		return i - a
	}
	// return func(i int) {
	// 	return i - a
	// }
}

//-----
func main() {
	// add2 := mkAdd(2)
	// add3 := mkAdd(3)
	// fmt.Println(add2(5), add3(6))

	var add1 = mkAdd(2)
	fmt.Println(add1(5))

	var add2 = mkAdd(3)
	fmt.Println(add2(6))

	//-----

	var add Add = func(i, j int) int { return i + j }
	add3 := add.Apply(3)
	fmt.Println("add 3 to 2:", add3(2))

	//-----

	in := 12          // imagine this gives us something
	m := multiply(4)  // example data
	s := subtract(10) // example data
	// subtract then multiply
	sm := func(i int) int { return m(s(i)) }
	ms := func(i int) int { return s(m(i)) }
	fmt.Printf("%v %v\n", ms(in), sm(in))
}
