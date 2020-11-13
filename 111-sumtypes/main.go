package main

import (
	"fmt"
)

//---

type ContactInfo interface {
	isContactInfo()
}

type EmailContactInfo struct{ name string }
type PostalContactInfo struct{ name string }

type BothContactMethods struct {
	EmailContactInfo
	PostalContactInfo
}

func (e EmailContactInfo) isContactInfo()   {}
func (p PostalContactInfo) isContactInfo()  {}
func (b BothContactMethods) isContactInfo() {}

// https://www.toptal.com/go/your-introductory-course-to-testing-with-go
// Make sure we can create variables of this type, w/o incurring any costs
// A great memory-free trick for ensuring that the interface is satisfied at
// run time is to insert the following into our code:
// This checks the assertion but doesn’t allocate anything, which lets us make
// sure that the interface is correctly implemented at compile time, before the
// program actually runs into any functionality using it.
var (
	_ ContactInfo = (*EmailContactInfo)(nil)
	_ ContactInfo = (*PostalContactInfo)(nil)
	_ ContactInfo = (*BothContactMethods)(nil)
)

func chooseContactInfo(c ContactInfo) ContactInfo {
	switch value := c.(type) {
	case EmailContactInfo:
		fmt.Printf("EmailContactInfo: %v", value)
		return value
	case PostalContactInfo:
		fmt.Printf("PostalContactInfo: %v", value)
		return value
	case BothContactMethods:
		fmt.Printf("BothContactMethods: %v", value)
		return value
	default:
		fmt.Printf("default: %v", value)
		panic("WRONG TYPE")
	}
}

//-----

type Contact struct {
	Name string
	ContactInfo
}

//-----

func main() {
	fmt.Println("---------------------------")
	emailcontactinfo := EmailContactInfo{name: "george@omame.com"}
	chooseContactInfo(emailcontactinfo)

	name := "George"

	contact := Contact{
		Name:        name,
		ContactInfo: emailcontactinfo,
	}
	fmt.Printf("%+v\n", contact)
}
