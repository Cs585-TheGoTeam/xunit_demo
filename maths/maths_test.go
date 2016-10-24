package maths

import "testing"

func TestAverage(t *testing.T) {
	var mean float64
	mean = Average([]float64{1,2})

	if mean != 1.5{
		t.Error("Expected 1.5, got ", mean)
	}
}