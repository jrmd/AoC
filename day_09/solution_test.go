package day_09

import (
	"testing"

	"github.com/jrmd/AoC/utils"
)

func GetInput() [][]int {
	lines, err := utils.GetLines("./test.txt")
	if err != nil {
		panic(err)
	}

	result, err := ParseFile(lines)
	if err != nil {
		panic(err)
	}

	return result
}

func TestPart1(t *testing.T) {
	input := GetInput()
	result := PartOne(input)

	if result != 114 {
		t.Fatalf("Expected 114, got %d", result)
	}
}

func TestPart2(t *testing.T) {
	input := GetInput()
	result := PartTwo(input)

	if result != 2 {
		t.Fatalf("Expected 2, got %d", result)
	}
}
