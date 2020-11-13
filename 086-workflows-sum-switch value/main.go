package main

import (
	"fmt"
)

type EnvelopeContents string

type CategorizedMail interface {
	isCategorizedMail()
}

type QuoteForm struct{ quantity float64 }
type OrderForm struct{ quantity float64 }

func (q QuoteForm) isCategorizedMail() {}
func (o OrderForm) isCategorizedMail() {}

var (
	c1 CategorizedMail = (*QuoteForm)(nil)
	c2 CategorizedMail = (*OrderForm)(nil)
)

func chooseCategorizedMail(c CategorizedMail) float64 {
	switch value := c.(type) {
	case QuoteForm:
		fmt.Printf("%#v\n", value)
		fmt.Printf("%#v\n", value.quantity*2)
		return value.quantity
	case OrderForm:
		fmt.Printf("%#v\n", value.quantity*10.1)
		return value.quantity
	default:
		panic(fmt.Sprintf("PANIC: %v", c))
	}
}

type Test struct{}

//----------------------------------------------
type ProductCatalog struct{ price float64 }
type PricedOrder float64

type CalculatePrices func(OrderForm, ProductCatalog) PricedOrder

var calculatePrices CalculatePrices = func(o OrderForm, p ProductCatalog) PricedOrder {
	total := o.quantity * p.price
	pricedorder := PricedOrder(total)
	return pricedorder
}

//----------------------------------------------

func main() {
	fmt.Println("----------------------------")

	category := OrderForm{quantity: 29.54}
	fmt.Printf("%#v\n", chooseCategorizedMail(category))

	category2 := QuoteForm{quantity: 1000.97}
	fmt.Printf("%#v\n", chooseCategorizedMail(category2))

	o := OrderForm{quantity: 99.3}
	p := ProductCatalog{price: 45.3}
	fmt.Println("calculatePrices: ", calculatePrices(o, p))

	// test := Test{}
	// chooseCategorizedMail(Test)
}
