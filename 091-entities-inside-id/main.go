/*
It’s more common to store the ID using the “inside” approach,
where each case has a copy of the identifier. Applied here,
we would create two separate types, one for each case
(UnpaidInvoice and PaidInvoice), both of which have their own
InvoiceId, and then a top-level Invoice type, which is a choice
between them.
*/
package main

import (
	"fmt"
	uuid "github.com/google/uuid"
)

type InvoiceId interface{}

type Invoice interface {
	isInvoice()
}

type UnpaidInvoice struct {
	InvoiceId
}

type PaidInvoice struct {
	InvoiceId
}

func (u UnpaidInvoice) isInvoice() {}
func (p PaidInvoice) isInvoice()   {}

var (
	_ Invoice = (*UnpaidInvoice)(nil)
	_ Invoice = (*PaidInvoice)(nil)
)

// The benefit of this approach is that now, when we do our
// pattern matching, we have all the data accessible in one place,
// including the ID:
func chooseInvoice(i Invoice) {
	switch value := i.(type) {
	case UnpaidInvoice:
		fmt.Printf("The unpaid invoiceid is %v\n", value.InvoiceId)
	case PaidInvoice:
		fmt.Printf("The paid invoiceis is %v\n", value.InvoiceId)
	default:
		panic("Wrong data")
	}
}

func main() {

	fmt.Println("-----------------------------")
	a := PaidInvoice{InvoiceId: InvoiceId(uuid.New())}
	chooseInvoice(a)

	b := UnpaidInvoice{InvoiceId: InvoiceId(uuid.New())}
	chooseInvoice(b)

	// This is rejected:
	// chooseInvoice(322342342342343233234)
}
