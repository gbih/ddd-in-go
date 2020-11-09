/*
Very sloppy!
Need to figure out the real meaning of Chp6, p.106 --
the real semantics -- and then use the appropriate idioms
and patterns.
*/

package unitquantity

type UnitQuantity struct {
	value int
}

// "smart constructor"
// func create(qty int) interface{} {
// 	return UnityQuantity{value: qty}
// }

func (u UnitQuantity) Create(qty int) int {
	if qty < 1 {
		panic("UnitQuantity can not be negative")
	}
	if qty > 1000 {
		panic("UnitQuantity can not be more than 1000")
	}
	return qty
}

// func Create(qty int) int {
// 	if qty < 1 {
// 		panic("UnitQuantity can not be negative")
// 	}
// 	if qty > 1000 {
// 		panic("UnitQuantity can not be more than 1000")
// 	}
// 	return qty
// }

// var Test = UnitQuantity{
// 	value: Create(132),
// }
// var UnitQtyResult = Test.value

var G UnitQuantity

func TestGB() int {
	return G.Create(343)
}

var UnitQtyResult = TestGB()
