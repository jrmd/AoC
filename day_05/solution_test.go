package day_05

import (
	"testing"

	"github.com/jrmd/AoC/utils"
)

func TestPart1(t *testing.T) {
	lines, err := utils.GetLines("./test.txt")
	if err != nil {
		panic(err)
	}

	data, err := ParseFile(lines)
	if err != nil {
		panic(err)
	}

	result := PartOne(data)

	if result != 35 {
		t.Fatalf("Expected 0, got %d", result)
	}
}

func TestPart2(t *testing.T) {
	lines, err := utils.GetLines("./test.txt")
	if err != nil {
		panic(err)
	}

	data, err := ParseFile(lines)

	if err != nil {
		panic(err)
	}

	result := PartTwo(data)

	if result != 46 {
		t.Fatalf("Expected 46, got %d", result)
	}
}
