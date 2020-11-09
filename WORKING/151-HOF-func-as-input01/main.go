/*
Mastering Go, Packt
Chapter 6, ps 228
What You Might Know About Go Packages
Functions that accept other functions as parameters
*/
package main

import "fmt"

// consistent signature of `func(i int) int`, a pseudo-internal API
func add2(i int) int {
	fmt.Println("add2() parameter i:", i)
	fmt.Printf("add2() %v + 2: \n", i)
	return i + 2
}

func sum3(i int) int {
	fmt.Println("sum3() parameter i:", i)
	return i * 3
}

// HOF function passed as parameter
// Here, we are just passing through, with no modification
// func funFun(f func(int) int, v int) int {
// 	fmt.Println("parameter v:", v)
// 	return f(v)
// }

func hof(f func(a int) int, b int) int {
	// a is undefined here, it is parameter of f only
	// b is directly accessible here, b is a parameter of hof
	fmt.Println("hof() parameter b:", b)
	fmt.Println("hof() intermediate step, f(b):", f(b))
	fmt.Println("hof() last step, f(b)+10:", f(b)+10)
	return f(b) + 10
}

//-----

func main() {
	fmt.Println("---------------------------")
	fmt.Println(hof(add2, 100))
	fmt.Println()
	fmt.Println(hof(add2, 200))
	fmt.Println()
	fmt.Println(hof(sum3, 10))
}
