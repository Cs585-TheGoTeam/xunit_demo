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

func TestMedian_Using_Odd_Length_Slice(t *testing.T) {
	slice := []float64{21,15,7,25,8}
	median := Median(slice)
	
	if median != 15{
		t.Error("Expected 15 but received", median)
	}
}

func TestMedian_Using_Even_Length_Slice(t *testing.T) {
	slice := []float64{21,15,7,25,8,16}
	median := Median(slice)
	
	if median != 15.5{
		t.Error("Expected 16.5 but received", median)
	}
}