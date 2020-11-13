// "Make illegal states unrepresentable." Go-style
package main

import (
	"fmt"
	vs "github.com/capturing-rules/verificationservice"
)

func chooseCustomerEmail(c vs.CustomerEmail) vs.CustomerEmail {
	switch value := c.(type) {
	case vs.Unverified:
		fmt.Printf("UNVERIFIED: %v\n", value)
		return value
	case vs.Verified:
		fmt.Printf("VERIFIED: %v\n", value)
		return value

	default:
		fmt.Printf("%v: %T\n", value, value)
		panic("Wrong type")
	}
}

func Check(email string) {
	checkverification := vs.CheckVerification(email)
	if checkverification != (vs.Verified{}) {
		verified := vs.Verified(checkverification)
		// Prevent direct assignment via `verified := vs.Verified{Email: "izumi@omame.com"}``
		chooseCustomerEmail(verified)
	} else {
		fmt.Printf("FAIL CHECK: %v\n", email)
	}
}

// The workflow that sends a password-reset message must take a
// VerifiedEmailAddress parameter as input rather than a normal email address.
type SendPasswordResetEmail func(vs.Verified)

var sendpasswordresetemail SendPasswordResetEmail = func(email vs.Verified) {
	// do logic here
	fmt.Printf("%T\n", email)
	fmt.Println("sendpasswordresetemail")
}

//-----

func main() {
	fmt.Println("-----------------------")

	unverified := vs.Unverified{Email: "george@omame.com"}
	chooseCustomerEmail(unverified)

	Check("izumi@omame.com")
	Check("jiko@omame.com")

	// very hacky
	checkverification := vs.CheckVerification("george@omame.com")
	verified := vs.Verified(checkverification)
	sendpasswordresetemail(verified)
}

/*
Notes:

F# code:

type Record = { Name: string; Age: int }

type Abc =
    | A
    | B of double
    | C of Record

The Abc type defined above can be either A, B, or C
(and nothing else).
In the case of A there is no "payload".
With B and C there is a payload that comes with it.

In Go, we can assume the tags will always come with a payload,
eg there is no distinction between tag and type.


Reference
https://stackoverflow.com/questions/28800672/how-to-add-new-methods-to-an-existing-type-in-go
*/
