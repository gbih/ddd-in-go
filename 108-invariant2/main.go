/*
Exploring lack of array/slice covariance in Go
https://stackoverflow.com/questions/3839335/any-sensible-solution-to-the-lack-of-array-slice-covariance-in-go
*/
package main

import "fmt"

type List interface {
	At(i int) interface{}
	Len() int
}

func printItems(header string, items List) {
	for i := 0; i < items.Len(); i++ {
		fmt.Print(items.At(i), " ")
	}
	fmt.Println()
}

type IntList []int
type FloatList []float64

func (il IntList) At(i int) interface{}   { return il[i] }
func (fl FloatList) At(i int) interface{} { return fl[i] }

func (il IntList) Len() int   { return len(il) }
func (fl FloatList) Len() int { return len(fl) }

func main() {
	var iarr = []int{1, 2, 3}
	var farr = []float64{1.2, 2.0, 3.0}
	printItems("Integer array:", IntList(iarr))
	printItems("Float array:", FloatList(farr))
}
