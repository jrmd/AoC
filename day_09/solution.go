package day_09

import (
	"fmt"
	"strings"

	"github.com/jrmd/AoC/utils"
)

func ParseFile(lines []string) ([][]int, error) {
	results := make([][]int, len(lines))

	for i, line := range lines {
		params, err := utils.SliceAtoi(strings.Fields(line))

		if err != nil {
			return nil, err
		}

		results[i] = params
	}

	return results, nil
}

func GetDifferences(seq []int) []int {
	diffs := make([]int, len(seq)-1)

	for i := 0; i < len(seq)-1; i++ {
		diffs[i] = seq[i+1] - seq[i]
	}

	return diffs
}

func AllMatchValue(seq []int, value int) bool {
	for _, v := range seq {
		if v != value {
			return false
		}
	}

	return true
}

func FindNextInSequence(sequence []int) int {
	var commonDifferences [][]int
	commonDifferences = append(commonDifferences, sequence)
	for {
		differences := GetDifferences(commonDifferences[len(commonDifferences)-1])

		if AllMatchValue(differences, 0) {
			break
		}

		commonDifferences = append(commonDifferences, differences)
	}

	for i := len(commonDifferences) - 1; i > 0; i-- {
		diff := commonDifferences[i]

		last := commonDifferences[i-1][len(commonDifferences[i-1])-1]
		commonDifferences[i-1] = append(commonDifferences[i-1], last+diff[len(diff)-1])
	}

	return commonDifferences[0][len(commonDifferences[0])-1]
}

func FindPrevInSequence(sequence []int) int {
	var commonDifferences [][]int
	commonDifferences = append(commonDifferences, sequence)
	for {
		differences := GetDifferences(commonDifferences[len(commonDifferences)-1])

		if AllMatchValue(differences, 0) {
			break
		}

		commonDifferences = append(commonDifferences, differences)
	}

	for i := len(commonDifferences) - 1; i > 0; i-- {
		diff := commonDifferences[i]

		last := commonDifferences[i-1][0]
		commonDifferences[i-1] = append([]int{last - diff[0]}, commonDifferences[i-1]...)
	}

	return commonDifferences[0][0]
}

func PartOne(input [][]int) int {
	total := 0

	for _, seq := range input {
		total += FindNextInSequence(seq)
	}

	return total
}

func PartTwo(input [][]int) int {
	total := 0

	for _, seq := range input {
		total += FindPrevInSequence(seq)
	}

	return total
}

func Run() {
	lines, err := utils.GetLines("./day_09/input.txt")
	if err != nil {
		panic(err)
	}

	result, err := ParseFile(lines)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Part One: %d\n", PartOne(result))
	fmt.Printf("Part Two: %d\n", PartTwo(result))
}
