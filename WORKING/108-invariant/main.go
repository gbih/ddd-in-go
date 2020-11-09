/*
Go has the following nillable types: pointers, functions, maps, interfaces, channel and slices
https://github.com/golang/go/issues/33078
*/

package main

import (
	"fmt"
)

//------

type NonEmptyList interface {
	nonEmptyList()
}

type Order struct {
	menu []int
}

func (o Order) nonEmptyList(n []int) {
	if len(n) <= 0 {
		panic("Empty1")
	}
	if n == nil {
		panic("Empty2")
	}
	fmt.Println()
	fmt.Println("n is ", n)
}

func main() {

	order := Order{
		menu: []int{1, 2, 3},
	}
	fmt.Printf("%#v", order)
	fmt.Printf("%v", order.menu)
	test := order.menu
	order.nonEmptyList(test)

	// fmt.Println("-------")
	// order2 := Order{}
	// test2 := order2.menu
	//order2.nonEmptyList(test2)

	//order.nonEmptyList(order{menu})
}

//type testType test

// type George struct {
// 	doesitwork test
// }

//var t George = &G1{test: 99}
//var t George = &G1{99}

// var t George = (*G1)(nil) // (*main.G1)(nil)

//var t George = &G1{} // &G1{}

/////

//type OrderLine []int

//var o OrderLine = (&OrderLine)

// func (n OrderLine) isNil() {
// 	if len(n) <= 0 {
// 		panic("Empty1")
// 	}
// 	if n == nil {
// 		panic("Empty2")
// 	}
// 	fmt.Println("n is ", n)
// }

// type Order struct {
// 	//OrderLines NonEmptyList
// 	OrderLines NonNil
// }

//var _ NonEmptyList = (&NonEmptyList)
//var t NonEmptyList = &NonEmptyList{}

// fmt.Printf("%#v", a)
//------

// var s1 []int         // nil slice
// s2 := []int{}        // non-nil, empty slice
// s3 := make([]int, 0) // non-nil, empty slice

// fmt.Printf("s1: len(s1) is %v, cap(s1) is %v, s1==nil is %v, s1[:] is %v, s1[:] == nil is %v \n", len(s1), cap(s1), s1 == nil, s1[:], s1[:] == nil)
// fmt.Println("s1: ", len(s1), cap(s1), s1 == nil, s1[:], s1[:] == nil)
// fmt.Println("s2: ", len(s2), cap(s2), s2 == nil, s2[:], s2[:] == nil)
// fmt.Println("s3: ", len(s3), cap(s3), s3 == nil, s3[:], s3[:] == nil)

// for range s1 {
// }
// for range s2 {
// }
// for range s3 {
// }

// fmt.Println("-------------------------------")
// //primes := []int{2, 3, 5, 7, 11, 13}
// primes := []int{}
// order := Order{OrderLines: OrderLine(primes)}

// fmt.Printf("%#v\n", order)
// fmt.Printf("%#v\n", &Order{})
