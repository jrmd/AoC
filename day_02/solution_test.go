package day_02

import (
	"testing"

	"github.com/jrmd/AoC/utils"
)

func GetGames() []*Game {
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
	lines := GetGames()
	result := PartOne(lines)

	if result != 8 {
		t.Fatalf("Expected 8, got %d", result)
	}
}

func TestPart2(t *testing.T) {
	lines := GetGames()
	result := PartTwo(lines)

	if result != 2286 {
		t.Fatalf("Expected 2286, got %d", result)
	}
}
