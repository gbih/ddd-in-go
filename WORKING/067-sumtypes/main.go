package main

import (
	"fmt"
)

type OrderQuantity interface {
	isOrderQuantity()
}

type UnitQuantity int
type KilogramQuantity float32

func (u UnitQuantity) isOrderQuantity()     {}
func (k KilogramQuantity) isOrderQuantity() {}

var (
	_ OrderQuantity = (*UnitQuantity)(nil)
	_ OrderQuantity = (*KilogramQuantity)(nil)
)

// printf formatting: https://golang.org/pkg/fmt/
func printQuantity(i OrderQuantity) {
	switch i.(type) {
	case UnitQuantity:
		fmt.Printf("%d units\n", i) // base 10 integer
	case KilogramQuantity:
		fmt.Printf("%1.2f kg\n", i) // float, width 1, precision 2
	}
}

//----------------------------------------

func main() {
	anOrderQtyInUnits := UnitQuantity(10)
	anOrderQtyInKg := KilogramQuantity(2.5)

	printQuantity(anOrderQtyInUnits)
	printQuantity(anOrderQtyInKg)
}
