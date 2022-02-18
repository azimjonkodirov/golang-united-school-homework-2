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

func CalcSquare(sideLen float64, sidesNum int) float64 {
	const PI= 3.14
	var x float64 = 2
	square:=math.Pow(sideLen,x)
	sqrt:=math.Sqrt(3)
	var res float64
	if sidesNum == 3{
		res = (sqrt/4)*square
	}else if sidesNum == 4{
     res=square
	}else if sidesNum ==0{
		res=PI*square
	}else{
    return 0
	}
	return  res
}

func main(){
fmt.Println(CalcSquare(10.0, 0))
}