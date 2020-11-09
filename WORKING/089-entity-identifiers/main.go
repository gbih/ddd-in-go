// https://pkg.go.dev/github.com/google/uuid?tab=doc
package main

import (
	"fmt"
	guuid "github.com/google/uuid"
)

//type ContactId guuid.UUID
type ContactId interface{}
type PhoneNumber int
type EmailAddress int

type Contact struct {
	ContactId
	PhoneNumber
	EmailAddress
}

func genUUID() {
	id := guuid.New()
	fmt.Printf("github.com/google/uuid:         %s\n", id.String())
}

func generateContacts() interface{} {
	var sum []Contact
	var contact Contact

	for i := 1; i < 5; i++ {
		contact = Contact{
			ContactId(guuid.New()),
			PhoneNumber(i),
			EmailAddress(i),
		}
		sum = append(sum, contact)
	}
	return sum
}

func main() {
	fmt.Printf("%+v\n", generateContacts())

}
