package main

import (
	"fmt"
	"unsafe"
)

//----------------------------------------------------

type CheckNumber int
type CardNumber string

//----------------------------------------------------

type CardType interface {
	isCardType()
}

type Visa string
type Mastercard string

func (Visa) isCardType()       {}
func (Mastercard) isCardType() {}

var (
	_ CardType = (*Visa)(nil)
	_ CardType = (*Mastercard)(nil)
)

//----------------------------------------------------

type CreditCardInfo struct {
	CardType
	CardNumber
}

//----------------------------------------------------

// This is no longer a simple “enum” because some of the choices have
// data associated with them: the Check case has a CheckNumber and the
// Card case has CreditCardInfo.

type PaymentMethod interface {
	isPaymentMethod()
}
type Cash string
type Check CheckNumber
type Card CreditCardInfo

func (c Cash) isPaymentMethod()  {}
func (c Check) isPaymentMethod() {}
func (c Card) isPaymentMethod()  {}

var (
	_ PaymentMethod = (*Cash)(nil)
	_ PaymentMethod = (*Check)(nil)
	_ PaymentMethod = (*Card)(nil)
)

//----------------------------------------------------

type PaymentAmount float32

type Currency interface {
	isCurrency()
}
type EUR string
type USD string

func (e EUR) isCurrency() {}
func (u USD) isCurrency() {}

var (
	_ Currency = (*EUR)(nil)
	_ Currency = (*USD)(nil)
)

//----------------------------------------------------

// Payment is the top-level type.
type Payment struct {
	Amount PaymentAmount
	Currency
	Method PaymentMethod
}

//----------------------------------------------------

// Given an UnpaidInvoice and then a Payment, we can create a PaidInvoice.
type UnpaidInvoice float64
type PaidInvoice float64

type PayInvoiceType func(a UnpaidInvoice, p Payment) PaidInvoice

var PayInvoice PayInvoiceType = func(a UnpaidInvoice, p Payment) PaidInvoice {
	fmt.Printf("a: %T\n", a)
	fmt.Printf("p: %T\n", p.Amount)
	paidInvoice := PaidInvoice(a)
	return paidInvoice
}

type Test func(a, b int) int

var TestFunc Test = func(a, b int) int {
	return a + b
}

//----------------------------------------------------

type ConvertPaymentCurrencyType func(p Payment, target Currency) Payment

var ConvertPaymentCurrency = func() {
	var payment Payment
	fmt.Printf("PAYMENT %#v\n", payment)
}

//----------------------------------------------------

func main() {

	var checknumber CheckNumber
	fmt.Printf("%T, %v bytes\n", checknumber, unsafe.Sizeof(checknumber))

	var cardnumber CardNumber
	fmt.Printf("%T, %v bytes\n", cardnumber, unsafe.Sizeof(cardnumber))

	//----------

	var cardtype CardType
	fmt.Printf("%T, %v bytes\n", cardtype, unsafe.Sizeof(cardtype))

	var visa Visa
	fmt.Printf("%T, %v bytes\n", visa, unsafe.Sizeof(visa))

	var mastercard Mastercard
	fmt.Printf("%T, %v bytes\n", mastercard, unsafe.Sizeof(mastercard))

	//----------

	var creditcardinfo CreditCardInfo
	fmt.Printf("%T, %v bytes\n", creditcardinfo, unsafe.Sizeof(creditcardinfo))

	//----------

	var paymentmethod PaymentMethod
	fmt.Printf("%T, %v bytes\n", paymentmethod, unsafe.Sizeof(paymentmethod))
	var cash Cash
	fmt.Printf("%T, %v bytes\n", cash, unsafe.Sizeof(cash))
	var check Check
	fmt.Printf("%T, %v bytes\n", check, unsafe.Sizeof(check))
	var card Card
	fmt.Printf("%T, %v bytes\n", card, unsafe.Sizeof(card))

	//----------

	var paymentamount PaymentAmount
	fmt.Printf("%T, %v bytes\n", paymentamount, unsafe.Sizeof(paymentamount))
	var currency Currency
	fmt.Printf("%T, %v bytes\n", currency, unsafe.Sizeof(currency))
	var eur EUR
	fmt.Printf("%T, %v bytes\n", eur, unsafe.Sizeof(eur))
	var usd USD
	fmt.Printf("%T, %v bytes\n", usd, unsafe.Sizeof(usd))

	//----------

	var payment Payment
	fmt.Printf("%T, %v bytes\n", payment, unsafe.Sizeof(payment))

	//----------

	var unpaidinvoice UnpaidInvoice
	fmt.Printf("%T, %v bytes\n", unpaidinvoice, unsafe.Sizeof(unpaidinvoice))
	var paidinvoice PaidInvoice
	fmt.Printf("%T, %v bytes\n", paidinvoice, unsafe.Sizeof(paidinvoice))
	var payinvoicetype PayInvoiceType
	fmt.Printf("%T, %v bytes\n", payinvoicetype, unsafe.Sizeof(payinvoicetype))

	fmt.Printf("%T, %v bytes\n", PayInvoice, unsafe.Sizeof(PayInvoice))

	//----------

	ConvertPaymentCurrency()
	fmt.Println(TestFunc(1, 2))

	var amount Payment
	amount = Payment{Amount: 34.3}

	fmt.Printf("%#v\n", amount)
	fmt.Println(PayInvoice(1000.0, amount))
}
