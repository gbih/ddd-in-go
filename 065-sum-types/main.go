package main

import (
	"fmt"
)

type AppleVariety interface {
	isAppleVariety()
}

type GoldenDelicious string
type GrannySmith string
type Fuji string

func (g GoldenDelicious) isAppleVariety() {}
func (gr GrannySmith) isAppleVariety()    {}
func (f Fuji) isAppleVariety()            {}

var (
	_ AppleVariety = (*GoldenDelicious)(nil)
	_ AppleVariety = (*GrannySmith)(nil)
	_ AppleVariety = (*Fuji)(nil)
)

//------------------------------------------------

type BananaVariety interface {
	isBananaVariety()
}

type Cavendish string
type GrosMichel string
type Manzano string

func (Cavendish) isBananaVariety()  {}
func (GrosMichel) isBananaVariety() {}
func (Manzano) isBananaVariety()    {}

var (
	_ BananaVariety = (*Cavendish)(nil)
	_ BananaVariety = (*GrosMichel)(nil)
	_ BananaVariety = (*Manzano)(nil)
)

//------------------------------------------------

type CherryVariety interface {
	isCherryVariety()
}

type Montmorency string
type Bing string

func (Montmorency) isCherryVariety() {}
func (Bing) isCherryVariety()        {}

var (
	_ CherryVariety = (*Montmorency)(nil)
	_ CherryVariety = (*Bing)(nil)
)

//================================================

type FruitSnack interface {
	isFruitSnack()
}

type Apple string
type Banana string
type Cherries string

func (a Apple) isAppleVariety() {}
func (a Apple) isFruitSnack()   {}

func (b Banana) isBananaVariety() {}
func (a Banana) isFruitSnack()    {}

func (c Cherries) isCherryVariety() {}
func (c Cherries) isFruitSnack()    {}

var (
	_ FruitSnack = (*Apple)(nil)
	_ FruitSnack = (*Banana)(nil)
	_ FruitSnack = (*Cherries)(nil)
)

func chooseFruitSnank(i FruitSnack) {
	switch i.(type) {
	case Apple:
		fmt.Println("SELECT: Apple")
	case Banana:
		fmt.Println("SELECT: Banana")
	case Cherries:
		fmt.Println("SELECT: Cherries")
	}
}

//-------------------------------

func main() {
	fmt.Println("------------------------------")
	food := Apple("apple")
	chooseFruitSnank(food)

}
