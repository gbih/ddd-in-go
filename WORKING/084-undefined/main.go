package main

import (
	"fmt"
)

// F#
// type CustomerInfo = Undefined
type CustomerInfo struct{}
type ShippingAddress struct{}
type BillingAddress struct{}
type OrderLine struct{}
type BillingAmount struct{}

type Order struct {
	CustomerInfo
	ShippingAddress
	BillingAddress
	OrderLine
	BillingAmount
}

func main() {
	var order Order
	fmt.Printf("%#v", order)
}
