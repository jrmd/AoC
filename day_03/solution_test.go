package day_03

import (
	"testing"

	"github.com/jrmd/AoC/utils"
)

func GetRows() []Row {
	lines, err := utils.GetLines("./test.txt")
	if err != nil {
		panic(err)
	}

	games, err := ParseFile(lines)
	if err != nil {
		panic(err)
	}

	return games
}

func TestPart1(t *testing.T) {
	lines := GetRows()
	result := PartOne(lines)

	if result != 4361 {
		t.Fatalf("Expected 4361, got %d", result)
	}
}

func TestPart2(t *testing.T) {
	lines := GetRows()
	result := PartTwo(lines)

	if result != 467835 {
		t.Fatalf("Expected 467835, got %d", result)
	}
}
