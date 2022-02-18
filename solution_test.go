package main

import
(
	"testing"
)
func TestSquare(t *testing.T){
	got:=Azimjon(10.0, 0)
	var want float64 =314
	if got != want{
		t.Errorf("got %f want %f", got, want)
	}
}