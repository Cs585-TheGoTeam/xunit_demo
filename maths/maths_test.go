package maths

import (
	"testing"
)

func TestMean(t *testing.T) {
	var mean float64
	mean = Mean([]float64{1,2})

	if mean != 1.5{
		t.Error("Expected 1.5, got", mean)
	}
}

func TestFibonacci_At_Index_10(t *testing.T) {
	seq := Fibonacci(10)
	
	if seq[9] != 34{
		t.Error("Expected 34 but got", seq[9])
	}
}

func TestFibonacci_Length_Of_Returned_Slice(t *testing.T) {
	seq := Fibonacci(10)
	
	if len(seq) != 10{
		t.Error("Expected a slice of length 10 but got", len(seq))
	}
}