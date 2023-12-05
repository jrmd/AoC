package day_05

import (
	"fmt"
	"strings"

	"github.com/jrmd/AoC/utils"
)

type Map struct {
	Source      int
	Destination int
	Range       int
}

type Category struct {
	Key  string
	Maps []Map
}

func (c *Category) GetDestination(source int) int {
	for _, m := range c.Maps {
		if m.Source <= source && source < m.Source+m.Range {
			return m.Destination + (source - m.Source)
		}
	}
	return source
}

type Almanac struct {
	Seeds      []int
	Categories []Category
}

func (a *Almanac) GetLocation(source int) int {
	src := source
	for _, cat := range a.Categories {
		src = cat.GetDestination(src)
	}
	return src
}

func ParseFile(lines []string) (*Almanac, error) {
	seeds, err := utils.SliceAtoi(strings.Fields(lines[0][7:]))

	if err != nil {
		return nil, err
	}

	almanac := Almanac{
		Seeds: seeds,
	}

	category := Category{}
	for _, line := range lines[2:] {
		if len(line) == 0 {
			almanac.Categories = append(almanac.Categories, category)
			category = Category{}
			continue
		}

		if strings.HasSuffix(line, " map:") {
			category.Key = line[:len(line)-5]
			continue
		}

		data, err := utils.SliceAtoi(strings.Fields(line))

		if err != nil {
			continue
		}

		category.Maps = append(category.Maps, Map{
			Destination: data[0],
			Source:      data[1],
			Range:       data[2],
		})
	}

	almanac.Categories = append(almanac.Categories, category)

	return &almanac, nil
}

func PartOne(almanac *Almanac) int {
	loc := -1

	for _, seed := range almanac.Seeds {
		newLoc := almanac.GetLocation(seed)

		if loc == -1 || newLoc < loc {
			loc = newLoc
		}
	}

	return loc
}

func PartTwo(almanac *Almanac) int {
	loc := -1

	for i := 0; i < len(almanac.Seeds); i += 2 {
		base := almanac.Seeds[i]
		count := almanac.Seeds[i+1]

		for j := 0; j < count; j++ {
			newLoc := almanac.GetLocation(base + j)

			if loc == -1 || newLoc < loc {
				loc = newLoc
			}
		}
	}

	return loc
}

func Run() {
	lines, err := utils.GetLines("./day_05/input.txt")
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
