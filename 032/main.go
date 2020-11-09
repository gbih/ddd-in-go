package main

import (
	"fmt"
)

type Order interface {
	Val() Order
}

type OrderLine struct {
	Product  string
	Quantity int
	Price    float32
}

type CustomerInfo struct{ LastName string }
type ShippingAddress struct{ Address string }
type BillingAddress struct{ Address string }
type OrderLines struct{ OrderLine }
type OrderArray []OrderLine
type AmountToBill struct{ Cost float32 }

type Order2 struct {
	CustomerInfo
	ShippingAddress
	BillingAddress
	OrderArray
	AmountToBill
}

// Implement method for type-checking
func (c CustomerInfo) Val() Order    { return c }
func (s ShippingAddress) Val() Order { return s }
func (b BillingAddress) Val() Order  { return b }
func (o OrderLines) Val() Order      { return o }

func (a AmountToBill) Val() Order { return a }

var (
	_ Order = (*CustomerInfo)(nil)
	_ Order = (*ShippingAddress)(nil)
	_ Order = (*BillingAddress)(nil)
	_ Order = (*OrderLines)(nil)
	_ Order = (*AmountToBill)(nil)
)

func do(i Order) string {
	switch v := i.(type) {
	case CustomerInfo:
		return fmt.Sprintf("type CustomerInfo: %v", v)
	case ShippingAddress:
		return fmt.Sprintf("type ShippingAddress: %v", v)
	case BillingAddress:
		return fmt.Sprintf("type BillingAddress: %v", v)
	case OrderLines:
		return fmt.Sprintf("type OrderLines: %v", v)
	case AmountToBill:
		return fmt.Sprintf("type AmountToBill: %v", v)
	default:
		return fmt.Sprintf("Unexpected type: %v", v)
	}
}

func main() {
	customerInfo := CustomerInfo{LastName: "Baptista"}
	// shippingAddess := ShippingAddress{Address: "123 Yuzawa"}
	// billingAddress := BillingAddress{Address: "123 Yuzawa"}

	//orderLines := OrderLines{OrderLine{Product: "TV", Quantity: 343, Price: 33.3}}

	var orderArray OrderArray
	// orderArray = []OrderLine{
	// 	OrderLine{Product: "TV", Quantity: 343, Price: 33.34},
	// 	OrderLine{Product: "Bagel", Quantity: 1003, Price: 2.13},
	// 	OrderLine{Product: "Bike", Quantity: 3, Price: 443.55},
	// }

	orderArray = []OrderLine{
		OrderLine{Product: "TV", Quantity: 343, Price: 33.34},
		OrderLine{Product: "Bagel", Quantity: 1003, Price: 2.13},
		OrderLine{Product: "Bike", Quantity: 3, Price: 443.55},
	}

	orderArray2 := OrderArray{OrderLine{Product: "TV", Quantity: 343, Price: 33.34},
		OrderLine{Product: "Bagel", Quantity: 1003, Price: 2.13},
		OrderLine{Product: "Bike", Quantity: 3, Price: 443.55}}

	fmt.Printf("1. orderArray = %T\n", orderArray)
	fmt.Printf("2. orderArray = %#v\n", orderArray)
	fmt.Println("--------------------")
	fmt.Printf("1. orderArray = %T\n", orderArray2)
	fmt.Printf("2. orderArray = %#v\n", orderArray2)

	fmt.Println(do(customerInfo))
}
