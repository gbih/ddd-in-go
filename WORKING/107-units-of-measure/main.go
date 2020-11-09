// standardized unit types to maintain compatibility
// Microsoft.FSharp.Data.UnitSystems.SI
// https://github.com/martinlindhe/unit
// Ugh, what a mess. The cleanest approach is the F# library!

package main

import (
	"fmt"
)

// Length represents a SI unit of length (in meters, m)
type Length Unit

type dpi Length
type inch Length
type px Length

var dpiValue dpi = dpi(150)

// Unit represents a unit
type Unit float64

// https://github.com/martinlindhe/unit/blob/master/length.go

const (
	Meter Length = 1e0
	Inch         = Meter * 0.0254
	Foot         = Inch * 12
)

// Meters returns the length in m
func (l Length) Meters() float64 {
	return float64(l)
}

func main() {
	fmt.Printf("%T\n", dpiValue)
	dpiValue := 22.33
	fmt.Printf("%v\n", dpiValue)
	dpiValue = "test"

	m := 1 * Meter
	fmt.Println(m)

	ft := 1 * Foot
	fmt.Println(ft)

	fmt.Println(ft.Meters(), "meters")

	fiveMeters := 5.0 * Meter
	fmt.Println(fiveMeters)
}
