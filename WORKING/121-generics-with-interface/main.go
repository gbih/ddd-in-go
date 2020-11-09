package main

import (
	"fmt"
	"sort"
)

func Test() {
	s := []int{5, 2, 6, 3, 1, 4} // unsorted
	sort.Ints(s)
	//sort.Sort(s)
	fmt.Println(s)
}

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%s: %d", p.Name, p.Age)
}

// ByAge implements sort.Interface based on the Age field
type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

func Test2() {
	people := []Person{
		{"Bob", 31},
		{"John", 42},
		{"Michael", 17},
		{"Jenny", 26},
	}

	// 1. Define a set of methods for the slice type,
	// as with ByAge, and call sort.Sort.
	sort.Sort(ByAge(people))
	fmt.Println(people)

	/*
			   // Sort sorts data.
			   // It makes one call to data.Len to determine n, and O(n*log(n)) calls to
			   // data.Less and data.Swap. The sort is not guaranteed to be stable.
			   func Sort(data Interface) {
			   	n := data.Len()
			   	quickSort(data, 0, n, maxDepth(n))
			   }


		func quickSort(data Interface, a, b, maxDepth int) {
			for b-a > 12 { // Use ShellSort for slices <= 12 elements
				if maxDepth == 0 {
					heapSort(data, a, b)
					return
				}
				maxDepth--
				mlo, mhi := doPivot(data, a, b)
				// Avoiding recursion on the larger subproblem guarantees
				// a stack depth of at most lg(b-a).
				if mlo-a < b-mhi {
					quickSort(data, a, mlo, maxDepth)
					a = mhi // i.e., quickSort(data, mhi, b)
				} else {
					quickSort(data, mhi, b, maxDepth)
					b = mlo // i.e., quickSort(data, a, mlo)
				}
			}
			if b-a > 1 {
				// Do ShellSort pass with gap 6
				// It could be written in this simplified form cause b-a <= 12
				for i := a + 6; i < b; i++ {
					if data.Less(i, i-6) {
						data.Swap(i, i-6)
					}
				}
				insertionSort(data, a, b)
			}
		}

	*/

}

//-------

func main() {
	Test()
	Test2()
}

/*

https://blog.golang.org/why-generics
https://golang.org/pkg/sort/

Go generic programming today
So how do people write this kind of code in Go?

In Go you can write a single function that works for different slice
types by using an interface type, and defining a method on the slice
types you want to pass in. That is how the standard library's sort.Sort
function works.

https://golang.org/src/sort/sort.go?s=5416:5441#L206

// Sort sorts data.
// It makes one call to data.Len to determine n, and O(n*log(n)) calls to
// data.Less and data.Swap. The sort is not guaranteed to be stable.
func Sort(data Interface) {
	n := data.Len()
	quickSort(data, 0, n, maxDepth(n))
}


In other words, interface types in Go are a form of generic programming.
They let us capture the common aspects of different types and express
them as methods. We can then write functions that use those interface
types, and those functions will work for any type that implements those
methods.

But this approach falls short of what we want. With interfaces you have
to write the methods yourself. It's awkward to have to define a named
type with a couple of methods just to reverse a slice. And the methods
you write are exactly the same for each slice type, so in a sense we've
just moved and condensed the duplicate code, we haven't eliminated it.
Although interfaces are a form of generics, they don’t give us everything
we want from generics.

A different way of using interfaces for generics, which could get around
the need to write the methods yourself, would be to have the language
define methods for some kinds of types. That isn't something the
language supports today, but, for example, the language could define
that every slice type has an Index method that returns an element.
But in order to use that method in practice it would have to return
an empty interface type, and then we lose all the benefits of static
typing. More subtly, there would be no way to define a generic function
that takes two different slices with the same element type, or that
takes a map of one element type and returns a slice of the same element
type. Go is a statically typed language because that makes it easier
to write large programs; we don’t want to lose the benefits of static
typing in order to gain the benefits of generics.

Another approach would be to write a generic Reverse function using
the reflect package, but that is so awkward to write and slow to run
that few people do that. That approach also requires explicit type
assertions and has no static type checking.

Or, you could write a code generator that takes a type and generates
a Reverse function for slices of that type. There are several code
generators out there that do just that. But this adds another step
to every package that needs Reverse, it complicates the build because
all the different copies have to be compiled, and fixing a bug in
the master source requires re-generating all the instances, some of
which may be in different projects entirely.

All these approaches are awkward enough that I think most people who
have to reverse a slice in Go just write the function for the specific
slice type that they need. Then they'll need to write test cases for
the function, to make sure they didn't make a simple mistake like
the one I made initially. And they'll need to run those tests routinely.

However we do it, it means a lot of extra work just for a function
that looks exactly the same except for the element type. It's not
that it can't be done. It clearly can be done, and Go programmers
are doing it. It's just that there ought to be a better way.

For a statically typed language like Go, that better way is generics.
What I wrote earlier is that generic programming enables the
representation of functions and data structures in a generic form,
 with types factored out. That's exactly what we want here.


*/

//-------

// https://yourbasic.org/golang/how-to-sort-in-go/
// Sort custom data structures
// Use the generic sort.Sort and sort.Stable functions
// They sort any collection that implements the sort.Interface interface

/*
package sort
// A type, typically a collection, that satisfies sort.Interface can be
// sorted by the routines in this package. The methods require that the
// elements of the collection be enumerated by an integer index.

type Interface interface {
	// Len is the number of elements in the collection.
	Len() int
	// Less reports whether the element with
	// index i should sort before the element with index j.
	Less(i, j int) bool
	// Swap swaps the elements with indexes i and j.
	Swap(i, j int)
}

For example:
// IntSlice attaches the methods of Interface to []int, sorting in increasing order.
type IntSlice []int
func (p IntSlice) Len() int           { return len(p) }
func (p IntSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p IntSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
// Sort is a convenience method.
func (p IntSlice) Sort() { Sort(p) }


// Float64Slice attaches the methods of Interface to []float64, sorting in increasing order
// (not-a-number values are treated as less than other values).
type Float64Slice []float64
func (p Float64Slice) Len() int           { return len(p) }
func (p Float64Slice) Less(i, j int) bool { return p[i] < p[j] || isNaN(p[i]) && !isNaN(p[j]) }
func (p Float64Slice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
// Sort is a convenience method.
func (p Float64Slice) Sort() { Sort(p) }


// StringSlice attaches the methods of Interface to []string, sorting in increasing order.
type StringSlice []string
func (p StringSlice) Len() int           { return len(p) }
func (p StringSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p StringSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
// Sort is a convenience method.
func (p StringSlice) Sort() { Sort(p) }

*/
