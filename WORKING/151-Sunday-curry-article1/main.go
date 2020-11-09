// https://blog.usejournal.com/function-currying-in-go-a88672d6ebcf
package main

import "fmt"

func greet(greeting, name string) string {
	return fmt.Sprintf("%v %v", greeting, name)
}

// Wrap the Greet function into another function taking
// one argument and returning a closure.
// By returning string here, the enclosed functions also
// have to return a string.
func prefixedGreet(p string) func(string) string {
	// return string
	return func(n string) string {
		return greet(p, n)
		//return fmt.Sprintf("Outer argument p: %v, inner argument n: %v", p, n)
	}
}

// 1st argument p is essentially a closure, accessed via `prefixedGreet2(9)`
// 2nd argument (func(int) string) is implemented in the body, and has access to the closure p
// 1st argument use: prefixedGreet2(999)
// 2nd argument use: prefixedGreet2(999)(23)
func prefixedGreet2(p int) func(int) string {
	// implement the second function
	return func(n int) string {
		//return greet(p, n)
		// Essentially have access to value p as a closure
		return fmt.Sprintf("Closure-outer argument p: %v, inner argument n: %v", p, n)
	}
}

var (
	// func(string) string
	eveningFn = prefixedGreet("Good evening") // first argument
	morningFn = prefixedGreet("Good morning")
)

//----------------------------------------

func main() {
	//fmt.Println("Style 1. ", greet("Good Morning", "Sam"))
	fmt.Println("Style 2. ", prefixedGreet("Good morning")("John"))
	fmt.Println("Style 3. ", eveningFn("Ricardo")) // 2nd argument
	// fmt.Println("Style 3. ", morningFn("Isabel"))
	fmt.Println("Style 2. ", prefixedGreet2(999)(12345))
}
