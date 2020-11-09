package main

import (
	"fmt"
	"time"
)

type Command struct {
	data      interface{}
	timeStamp time.Time
	userId    string
}

type CommandData interface {
	Data() interface{}
}

func (c Command) Data(u interface{}) interface{} { return u }

func (c Command) Get() interface{} {
	return c.Data(c.Data(c))
}

//-----

func main() {
	fmt.Println("----------------------------")
	var c Command
	c = Command{
		data:      c.Data(999.54),
		timeStamp: time.Now(),
		userId:    "ID123",
	}
	//fmt.Printf("c.Data(): %#v\n", c.Data(999.54))
	//c.data = c.Data(999.54)
	fmt.Printf("TEST%#v\n", c)
	fmt.Println()
	fmt.Printf("c.Get(): %#v\n", c.Get())

	fmt.Printf("000: %#v\n", c)
	fmt.Printf("001: %#v\n", c.data)
	fmt.Printf("002: %#v\n", c.timeStamp)
	fmt.Printf("003: %#v\n", c.userId)

}

/*
Interface types in Go are a form of generic programming.
They let us capture the common aspects of different types and express
them as methods. We can then write functions that use those interface
types, and those functions will work for any type that implements those
methods.
*/
/*
F#:
type Command<'data> = {
	Data : 'data
	Timestamp: DateTime
	UserId: string
}

type PlaceOrder = Command<UnvalidatedOrder>

Go: Model after std-library package sort.Interface:
type Interface interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}
*/
