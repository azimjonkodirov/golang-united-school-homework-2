package main

import (
	"math"
	"testing"
)
func TestSquare(t *testing.T){
	got:=CalcSquare(10.0, 0)
	var want float64 =math.Pi*100.0
	if got != want{
		t.Errorf("got %f want %f", got, want)
	}
}