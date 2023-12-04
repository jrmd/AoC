package day_03

import (
	"fmt"
	"strconv"
	"unicode"

	"github.com/jrmd/AoC/utils"
)

func ParseFile(lines []string) ([]Row, error) {
	rows := []Row{}

	for _, line := range lines {
		row := Row{
			Numbers: []Number{},
		}

		start, end := -1, -1

		for i, ch := range line {
			if unicode.IsDigit(rune(ch)) {
				if start == -1 {
					start = i
				}
				end = i
				continue
			}

			if end > -1 {
				value, err := strconv.Atoi(line[start : end+1])
				if err != nil {
					return nil, err
				}
				row.Numbers = append(row.Numbers, Number{
					Value:         value,
					StartPosition: start,
					EndPosition:   end,
				})
			}

			start, end = -1, -1

			if ch == rune('.') {
				continue
			}

			if ch == rune('*') {
				row.Gears = append(row.Gears, Gear{
					Position: i,
				})
			}

			row.Symbols = append(row.Symbols, i)
		}

		if end > -1 {
			value, err := strconv.Atoi(line[start : end+1])
			if err != nil {
				return nil, err
			}
			row.Numbers = append(row.Numbers, Number{
				Value:         value,
				StartPosition: start,
				EndPosition:   end,
			})
		}

		rows = append(rows, row)
	}

	return rows, nil
}

type Number struct {
	Value         int
	StartPosition int
	EndPosition   int
}

type Gear struct {
	Position int
	Parts    []int
}

type Row struct {
	Numbers []Number
	Symbols []int
	Gears   []Gear
}

func (r *Row) HasSymbol(start int, end int) bool {
	for _, v := range r.Symbols {
		if v >= start && v <= end {
			return true
		}

		if v > end {
			break
		}
	}
	return false
}

func (r *Row) CheckGear(start int, end int, value int) {
	for i, v := range r.Gears {
		if v.Position >= start && v.Position <= end {
			r.Gears[i].Parts = append(r.Gears[i].Parts, value)
		}

		if v.Position > end {
			break
		}
	}
}

func PartOne(rows []Row) int {
	total := 0

	for i, row := range rows {
		for _, number := range row.Numbers {
			if i > 0 {
				if rows[i-1].HasSymbol(number.StartPosition-1, number.EndPosition+1) {
					total += number.Value
					continue
				}
			}

			if row.HasSymbol(number.StartPosition-1, number.EndPosition+1) {
				total += number.Value
				continue
			}

			if i < len(rows)-1 {
				if rows[i+1].HasSymbol(number.StartPosition-1, number.EndPosition+1) {
					total += number.Value
					continue
				}
			}
		}
	}

	return total
}

func PartTwo(rows []Row) int {
	total := 0

	for i, row := range rows {
		for _, number := range row.Numbers {
			if i > 0 {
				rows[i-1].CheckGear(number.StartPosition-1, number.EndPosition+1, number.Value)
			}

			row.CheckGear(number.StartPosition-1, number.EndPosition+1, number.Value)

			if i < len(rows)-1 {
				rows[i+1].CheckGear(number.StartPosition-1, number.EndPosition+1, number.Value)
			}
		}
	}

	for _, row := range rows {
		for _, gear := range row.Gears {
			if len(gear.Parts) == 2 {
				total += gear.Parts[0] * gear.Parts[1]
			}
		}
	}

	return total
}

func Run() {
	lines, err := utils.GetLines("./day_03/input.txt")
	if err != nil {
		panic(err)
	}

	rows, err := ParseFile(lines)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Part One: %d\n", PartOne(rows))
	fmt.Printf("Part Two: %d\n", PartTwo(rows))
}
