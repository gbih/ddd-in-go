package main

import (
	"fmt"
)

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
	fmt.Printf("%#v\n", order)
}
