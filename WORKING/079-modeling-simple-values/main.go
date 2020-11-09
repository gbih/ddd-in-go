package main

import (
	"fmt"
)

// This is a reference type
type CustomerID struct {
	float64
}

// This is a value type, but using it as a pointer prevents casting back to float64
type CustomerID2 float64
type CustomerID3 float64

// Go does not permit casting a pointer to another pointer
// This is how we can prevent passing simple float64 into here
func TestReference(input *CustomerID2) CustomerID2 {
	fmt.Printf("input is %T\n", input)
	return *input
}

func TestStruct(input CustomerID) CustomerID {
	fmt.Printf("input is %#v\n", input)
	return input
}

///

func main() {
	customerid2 := CustomerID2(33.33)
	fmt.Println(TestReference(&customerid2))

	customerid2b := CustomerID2(133.33)

	//customerid3 := CustomerID3(133.33)

	fmt.Println(customerid2 == customerid2b)

	//fmt.Println(customerid2 == customerid3)
	//fmt.Println(TestReference(&customerid3))

	// b := CustomerID3(55.44)
	// fmt.Println(TestReference(&b))

	// fmt.Println(TestReference(334.222))

	c := CustomerID{3343.343}
	fmt.Println(TestStruct(c))
	fmt.Printf("%+v", TestStruct(c))
}
