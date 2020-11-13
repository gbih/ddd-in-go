package main

import (
	"fmt"
)

type AcknowledgmentSent interface{}
type OrderPlaced interface{}
type BillableOrderPlaced interface{}
type UnvalidatedOrder interface{}

type PlaceOrderEvents struct {
	AcknowledgmentSent
	OrderPlaced
	BillableOrderPlaced
}

type PlaceOrder func(UnvalidatedOrder) PlaceOrderEvents

var placeOrder PlaceOrder = func(u UnvalidatedOrder) PlaceOrderEvents {
	result := PlaceOrderEvents{
		AcknowledgmentSent:  1,
		OrderPlaced:         u,
		BillableOrderPlaced: 1.2,
	}
	return result
}

func main() {
	fmt.Printf("%+v\n", placeOrder("Pizza"))
}
