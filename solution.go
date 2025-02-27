package main

import (
	"fmt"
	"math"
)

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

type sides int

const SidesTriangle sides = 3
const SidesSquare sides = 4
const SidesCircle sides = 0

func CalcSquare(sideLen float64, sidesNum sides) (res float64) {
	switch sidesNum {
	case SidesCircle:
		res = math.Pi * math.Pow(sideLen, 2.0)
	case SidesTriangle:
		res = (math.Pow(sideLen, 2.0) * math.Sqrt(3.0)) / 4
	case SidesSquare:
		res = math.Pow(sideLen, 2.0)
	default:
		res = 0
	}

	return
}

func main(){
fmt.Println(CalcSquare(10.0, 0))
}