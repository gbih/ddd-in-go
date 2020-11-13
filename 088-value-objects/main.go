package main

import "fmt"
import vo "github.com/gbih/ddd-in-go/088-value-objects/immutable/valueobject001"

func main() {
	// Record types
	fmt.Printf("%#v\n", vo.Name1)
	fmt.Println(vo.Name1 == vo.Name2)
	fmt.Println(vo.Name1 == vo.Name3)

	// This mutation is not possible:
	// vo.Name1.firstname = "Michael"

	fmt.Println("------------")
	fmt.Println(vo.Address1 == vo.Address2)

	fmt.Println("------------")
	fmt.Println(vo.Phone1)
	fmt.Println(vo.Phone1 == vo.Phone2)
	fmt.Println(vo.Phone1 == vo.Phone3)

	// Choice types (sum types)
	// cannot compare different types
	// fmt.Println(vo.Phone1 == vo.Order4)
}
