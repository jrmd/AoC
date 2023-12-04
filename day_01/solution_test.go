package day_01

import (
	"testing"
)

func TestPart1(t *testing.T) {
	lines := []string{
		"1abc2",
		"pqr3stu8vwx",
		"a1b2c3d4e5f",
		"treb7uchet",
	}

	result := PartOne(lines)

	if result != 142 {
		t.Fatalf("Expected 142, got %d", result)
	}
}

func TestPart2(t *testing.T) {
	lines := []string{
		"two1nine",
		"eightwothree",
		"abcone2threexyz",
		"xtwone3four",
		"4nineeightseven2",
		"zoneight234",
		"7pqrstsixteen",
	}

	result := PartTwo(lines)

	if result != 281 {
		t.Fatalf("Expected 281, got %d", result)
	}
}
