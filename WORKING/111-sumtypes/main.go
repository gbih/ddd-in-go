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

// make sure we can create variables of this type, w/o incurring any costs
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
