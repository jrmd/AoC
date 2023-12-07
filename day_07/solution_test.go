package day_07

import (
	"testing"

	"github.com/jrmd/AoC/utils"
)

func GetInput() []string {
	lines, err := utils.GetLines("./test.txt")
	if err != nil {
		panic(err)
	}

	return lines
}

func TestPart1(t *testing.T) {
	games := GetInput()
	result := PartOne(games)

	if result != 6440 {
		t.Fatalf("Expected 6440, got %d", result)
	}
}

func TestPart2(t *testing.T) {

	games := GetInput()
	result := PartTwo(games)

	if result != 5905 {
		t.Fatalf("Expected 5905, got %d", result)
	}
}
