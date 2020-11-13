/*
Go has the following nillable types: pointers, functions, maps, interfaces, channel and slices
https://github.com/golang/go/issues/33078
*/

package main

import (
	"fmt"
)

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

}
