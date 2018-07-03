package main

import (
	"testing"
	"math"
)

const float64EqualityThreshold = 1e-4

func almostEqual(a, b float64) bool {
	return math.Abs(a - b) <= float64EqualityThreshold
}


func TestRpnCalculator(t *testing.T) {
	calc := New()
	err := calc.evaluate("19 2.14 + 4.5 2 4.3 / - *")
	if err != nil {
		t.Error(err)
	}
	top, ok := calc.top()
	if !ok {
		t.Errorf("stack should have at least 1 element")
	}
	if !almostEqual(top, 85.2974) {
		t.Errorf("top should be 85.2974, is %v", top)
	}
}
