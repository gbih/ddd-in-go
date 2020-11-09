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

// VERY DIRTY, NEED TO ORGANIZE
// Explicitly document that the workflow that sends a password-reset
// message must take a VerifiedEmailAddress parameter as input rather
// than a normal email address.
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

	// VERY DIRTY, NEED TO ORGANIZE
	checkverification := vs.CheckVerification("george@omame.com")
	verified := vs.Verified(checkverification)
	sendpasswordresetemail(verified)
}

/*
Notes:

Init module:
go mod init github.com/capturing-rules

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


*/

/*
https://stackoverflow.com/questions/28800672/how-to-add-new-methods-to-an-existing-type-in-go

126

I wanted to expand on the answer given by @jimt here. That answer is correct and helped me tremendously in sorting this out. However, there are some caveats to both methods (alias, embed) with which I had trouble.

note: I use the terms parent and child, though I'm not sure that is the best for composition. Basically, parent is the type which you want to modify locally. Child is the new type that attempts to implement that modification.

Method 1 - Type Definition
type child parent
// or
type MyThing imported.Thing
Provides access to the fields.
Does not provide access to the methods.
Method 2 - Embedding (official documentation)
type child struct {
    parent
}
// or with import and pointer
type MyThing struct {
    *imported.Thing
}
Provides access to the fields.
Provides access to the methods.
Requires consideration for initialization.
Summary
Using the composition method the embedded parent will not initialize if it is a pointer. The parent must be initialized separately.
If the embedded parent is a pointer and is not initialized when the child is initialized, a nil pointer dereference error will occur.
Both type definition and embed cases provide access to the fields of the parent.
The type definition does not allow access to the parent's methods, but embedding the parent does.
You can see this in the following code.


*/
