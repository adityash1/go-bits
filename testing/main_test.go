package main

import (
	"testing"
)

func TestCalculateValues(t *testing.T) {
	var (
		expected = 10
		a        = 5
		b        = 5
	)
	have := calculateValues(a, b)
	if have != expected {
		t.Errorf("expected %d but got %d", expected, have)
	}

}
