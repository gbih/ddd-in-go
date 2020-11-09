/*
Modeling an Order as a Set of States
steps in the workflow pipeline.
The Order workflow  isn’t just a static document but
transitions through a series of different states
Unprocesses order form
Unvalidated order
Validated order
Priced order
Invalid order
*/
package main

import (
	"fmt"
)

type OrderId string
type CustomerInfo string
type ShippingAddress string
type BillingAddress string
type OrderLines string

type ValidatedOrder struct {
	OrderId
	CustomerInfo
	ShippingAddress
	BillingAddress
	OrderLines
}

//---

type AmountToBill string

type PricedOrder struct {
	OrderId
	CustomerInfo
	ShippingAddress
	BillingAddress
	OrderLines
	AmountToBill
}

//---

// This is the object that represents the order at any
// time in its life cycle. And this is the type that
// can be persisted to storage or communicated to other
// contexts.

type Order interface {
	isOrder()
}

type Unvalidated struct{}
type Validated struct{}
type Priced struct{}

func (u Unvalidated) isOrder() {}
func (v Validated) isOrder()   {}
func (p Priced) isOrder()      {}

var (
	_ Order = (*Unvalidated)(nil)
	_ Order = (*Validated)(nil)
	_ Order = (*Priced)(nil)
)

//---
func main() {
	fmt.Println()
}

/*
A much better approach is to make each state have its own type, which stores the data that is relevant to that state (if any). The entire set of states can then be represented by a choice type with a case for each state.
*/
