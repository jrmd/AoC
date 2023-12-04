package day_04

import (
	"testing"

	"github.com/jrmd/AoC/utils"
)

func TestPart1(t *testing.T) {
	lines, err := utils.GetLines("./test.txt")
	if err != nil {
		panic(err)
	}

	games, err := ParseFile(lines)
	if err != nil {
		panic(err)
	}

	result := PartOne(games)

	if result != 13 {
		t.Fatalf("Expected 13, got %d", result)
	}
}

func TestPart2(t *testing.T) {
	lines, err := utils.GetLines("./test.txt")
	if err != nil {
		panic(err)
	}

	games, err := ParseFile(lines)
	if err != nil {
		panic(err)
	}
	result := PartTwo(games)

	if result != 30 {
		t.Fatalf("Expected 30, got %d", result)
	}
}
