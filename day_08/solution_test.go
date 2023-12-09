package day_08

import (
	"testing"

	"github.com/jrmd/AoC/utils"
)

func GetInput(file string) *Network {
	lines, err := utils.GetLines(file)
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
	network := GetInput("./test.txt")
	result := PartOne(network)

	if result != 6 {
		t.Fatalf("Expected 6, got %d", result)
	}
}

func TestPart2(t *testing.T) {
	network := GetInput("./test-2.txt")
	result := PartTwo(network)

	if result != 6 {
		t.Fatalf("Expected 6, got %d", result)
	}
}
