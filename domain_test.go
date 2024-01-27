package main

import (
	"testing"
)

func TestCalculatePoints(t *testing.T) {
	receipt := Receipt{Total: "100.00"}

	expectedPoints := 5

	receipt.calculatePoints()
	
	if receipt.Points != expectedPoints {
		t.Errorf("calculatePoints() = %v; want %v", points, expectedPoints)
	}
}
