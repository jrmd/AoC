package day_01

import (
	"testing"

	"github.com/jrmd/AoC/utils"
)

func GetInput(file string) (IntSlice, IntSlice) {
	lines, err := utils.GetLines(file)
	if err != nil {
		panic(err)
	}

	lhs, rhs, err := ParseFile(lines)
	if err != nil {
		panic(err)
	}

	return lhs, rhs
}

func TestPart1(t *testing.T) {
  lhs, rhs := GetInput("./test.txt")
	result := PartOne(lhs, rhs)

	if result != 11 {
		t.Fatalf("Expected 11, got %d", result)
	}
}

func TestPart2(t *testing.T) {
  lhs, rhs := GetInput("./test.txt")
	result := PartTwo(lhs, rhs)

	if result != 31 {
		t.Fatalf("Expected 31, got %d", result)
	}
}
