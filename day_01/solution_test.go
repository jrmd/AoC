package template

import (
	"testing"
)

func TestPart1(t *testing.T) {
	result := PartOne()

	if result != 0 {
		t.Fatalf("Expected 0, got %d", result)
	}
}

func TestPart2(t *testing.T) {
	result := PartTwo()

	if result != 0 {
		t.Fatalf("Expected 0, got %d", result)
	}
}
