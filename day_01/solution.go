package day_01

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"github.com/jrmd/AoC/utils"
)

func GetIntFromString(line string, extract func(input string) (string, bool)) int {
	var firstInt, secondInt string

	for i := 0; i < len(line); i++ {
		if result, ok := extract(line[i:]); ok {
			firstInt = result
			break
		}
	}

	for i := len(line) - 1; i >= 0; i-- {
		if result, ok := extract(line[i:]); ok {
			secondInt = result
			break
		}
	}

	val, err := strconv.Atoi(firstInt + secondInt)
	if err != nil {
		return 0
	}

	return val
}

func PartOne(lines []string) int {
	total := 0

	for _, line := range lines {
		total += GetIntFromString(line, func(input string) (string, bool) {
			if unicode.IsDigit(rune(input[0])) {
				return string(input[0]), true
			}
			return "", false
		})
	}
	return total
}

func PartTwo(lines []string) int {
	total := 0

	for _, line := range lines {
		slice := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
		total += GetIntFromString(line, func(input string) (string, bool) {
			if unicode.IsDigit(rune(input[0])) {
				return string(input[0]), true
			}

			for i, key := range slice {
				if strings.HasPrefix(input, key) {
					return strconv.Itoa(i + 1), true
				}
			}

			return "", false
		})
	}

	return total
}

func Run() {
	lines, err := utils.GetLines("./day_01/input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Part One: %d\n", PartOne(lines))
	fmt.Printf("Part Two: %d\n", PartTwo(lines))
}
