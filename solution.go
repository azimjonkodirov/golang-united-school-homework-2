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

func CalcSquare(sideLen float64, sidesNum sides) float64 {
	const PI= 3.14
	var x float64 = 2
	square:=math.Pow(sideLen,x)
	sqrt:=math.Sqrt(3)
	
	var result float64
	switch sidesNum {
	case SidesCircle:
		result = PI * square
	case SidesTriangle:
		result = (sqrt * square) / 4
	case SidesSquare:
		result = square
	default:
		result = 0
	}

	return result
}

func main(){
fmt.Println(CalcSquare(10.0, 0))
}