package day_06

import (
	"testing"

	"github.com/jrmd/AoC/utils"
)

func GetInput() []Race {
	lines, err := utils.GetLines("./test.txt")
	if err != nil {
		panic(err)
	}

	races, err := ParseFile(lines)
	if err != nil {
		panic(err)
	}

	return races
}

func TestPart1(t *testing.T) {
	result := PartOne(GetInput())

	if result != 288 {
		t.Fatalf("Expected 288, got %d", result)
	}
}

func TestPart2(t *testing.T) {
	lines, err := utils.GetLines("./test.txt")
	if err != nil {
		panic(err)
	}
	result := PartTwo(lines)

	if result != 71503 {
		t.Fatalf("Expected 71503, got %d", result)
	}
}
