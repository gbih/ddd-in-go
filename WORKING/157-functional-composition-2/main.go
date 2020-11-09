package main

// http://www.owlfish.com/software/go/functionalComposition.html

import "fmt"

// F#
// let inc x = x + 1
// let sqr x = x * x
// let result = (8 |> sqr) |> inc  // 65

/*
Micro languages
Functional composition can also be used to create micro languages that
enable complex behaviour to be built up. For example, parsing a CSV file
results in a series of []string records. If we wish to filter out particular
records of interest we could write lots of nested if statements, but the
logic can become complex and error prone.

As an alternative we can create a series of small composable functions
that allow the logic of the filtering to be constructed more clearly:
*/
type Filter func(record []string) bool

// GB: The key to functional composition here is these functions all share the
// same signature, which allows them to be used as input and output and
// chained together.

// The Equal function returns a Filter function that takes []string records
// and returns true if the given column matches the specified value.
func Equal(column int, value string) Filter {
	return func(record []string) bool {
		if len(record) <= column {
			return false
		}
		if record[column] == value {
			return true
		}
		return false
	}
}

// And and Or functions can be created that string filters together.
func And(filters ...Filter) Filter {
	return func(record []string) bool {
		for _, f := range filters {
			if !f(record) {
				return false
			}
		}
		return true
	}
}

func Or(filters ...Filter) Filter {
	return func(record []string) bool {
		for _, f := range filters {
			if f(record) {
				return true
			}
		}
		return false
	}
}

// These allow us to build moderately complex filtering logic in a readable
// and easily extendable way:

func main() {
	testRecord := []string{"apple", "pear", "orange", "bear"}
	query1 := And(Equal(0, "apple"), Equal(1, "lemon"))
	query2 := Or(query1, And(Equal(2, "orange"), Equal(3, "bear")))

	if query1(testRecord) {
		fmt.Println("Query 1 will not match")
	}
	if query2(testRecord) {
		fmt.Println("Query2 matched")
	}
	// Query2 matched
}
