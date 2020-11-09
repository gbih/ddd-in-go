package main

import (
	"fmt"
	"time"
)

// F#
// type Command<'data> = {
// 		Data: 'data
//		Timestamp: DateTime
// 		UserId: string
// }


// Go 2.0 with Generics
// type Command[T any] struct {
// 		Data   T
//		TimeStamp time.Time
// 		UserId string
// }


// Go 1.15
// Essentially using Sum-type to constrain data type. 
// We have to declare beforehand, but this allows type association,
// for a pseudo-generic behavior
type Command struct {
	data      Data
	timeStamp time.Time
	userId    string
}

//---

type Data interface {
	isDataType()
}

type UnvalidatedOrder struct{ name string }
type ValidatedOrder struct{ name string }

func (u UnvalidatedOrder) isDataType() {}
func (v ValidatedOrder) isDataType()   {}

var (
	_ Data = (*UnvalidatedOrder)(nil)
	_ Data = (*ValidatedOrder)(nil)
)

//---

func Debug(input Command) {
	fmt.Println("=========")
	fmt.Println("FUNC Debug():")
	fmt.Printf("001: %#v\n", input.data)
	fmt.Printf("002: %#v\n", input.timeStamp)
	fmt.Printf("003: %#v\n", input.userId)
	fmt.Println("=========")
}

// func checkData(d Data) Data {
// 	switch value := d.(type) {
// 	case UnvalidatedOrder:
// 		fmt.Println("checkData", value.name)
// 		return d
// 	case ValidatedOrder:
// 		fmt.Println("checkData", value.name)
// 		return d
// 	default:
// 		panic("WRONG TYPE")
// 	}
// }

func (c Command) checkData(d Data) Data {
	switch value := d.(type) {
	case UnvalidatedOrder:
		fmt.Println("UnvalidatedOrder:", value.name)
		return d
	case ValidatedOrder:
		fmt.Println("ValidatedOrder:", value.name)
		return d
	default:
		panic("WRONG TYPE")
	}
}

//-----

func main() {
	fmt.Println("----------------------------")

	b := Command{
		data:      UnvalidatedOrder{name: "12B"},
		timeStamp: time.Now(),
		userId:    "ID-123",
	}
	Debug(b)
	b.checkData(b.data)

	//---

	c := Command{
		data:      ValidatedOrder{name: "23433j"},
		timeStamp: time.Now(),
		userId:    "ID-123",
	}

	fmt.Println()
	Debug(c)
	c.checkData(c.data)
	fmt.Println()

	//---

	d := Command{
		//data:      123, // not possible!
		data:      nil,
		timeStamp: time.Now(),
		userId:    "ID-789",
	}
	fmt.Println()
	Debug(d)
	d.checkData(d.data)
	fmt.Println()
}
