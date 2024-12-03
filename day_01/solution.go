package day_01

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/jrmd/AoC/utils"
)

type IntSlice []int

func (s IntSlice) Len() int {
   return len(s)
}

func (s IntSlice) Swap(i, j int) {
   s[i], s[j] = s[j], s[i]
}

func (s IntSlice) Less(i, j int) bool {
   return s[i] > s[j]
}

func Distance(a, b int) int {
   if a < b {
      return b - a
   }
   return a - b
}

func ParseFile(lines []string) (IntSlice, IntSlice, error) {
  leftSide := IntSlice{}
  rightSide := IntSlice{}

  for _, line := range lines {
    vals := strings.Split(line, "   ");

    lhi, err := strconv.Atoi(vals[0])
    rhi, err := strconv.Atoi(vals[1])
		if err != nil {
      panic(err)
		}

    leftSide = append(leftSide, lhi)
    rightSide = append(rightSide, rhi)
  }

  sort.Sort(leftSide)
  sort.Sort(rightSide)

	return leftSide, rightSide, nil
}

func PartOne(leftSide, rightSide IntSlice) int {
  total := 0
  for i := range leftSide {
    total += Distance(leftSide[i], rightSide[i])
  }
	return total;
}

func PartTwo(lhs, rhs IntSlice) int {
  rhsMap := map[int]int{}
  for _, val := range rhs {
    _, ok := rhsMap[val];

    if !ok {
      rhsMap[val] = 1
      continue;
    }

    rhsMap[val] += 1
  }

  total := 0;

  for _, val := range lhs {
    occurences, ok := rhsMap[val];

    if !ok {
      continue;
    }

    total += val * occurences;
  }
	return total
}

func Run() {
	lines, err := utils.GetLines("./day_01/input.txt")
	if err != nil {
		panic(err)
	}

  leftSide, rightSide, err := ParseFile(lines)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Part One: %d\n", PartOne(leftSide, rightSide))
	fmt.Printf("Part Two: %d\n", PartTwo(leftSide, rightSide))
}
