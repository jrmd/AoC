package day_06

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jrmd/AoC/utils"
)

type Race struct {
	Time   int
	Record int
}

func (r *Race) GetMinSpeed() int {
	speed := r.Record / r.Time
	speed += 1

	for i := speed; i < r.Time; i++ {
		if (r.Time-i)*i > r.Record {
			return i
		}
	}

	return speed
}

func (r *Race) GetMaxSpeed(min int) int {
	for i := r.Time - 1; i >= min; i-- {
		if (r.Time-i)*i > r.Record {
			return i
		}
	}

	return 1
}

func ParseFile(lines []string) ([]Race, error) {
	times, err := utils.SliceAtoi(strings.Fields(lines[0][5:]))
	if err != nil {
		return nil, err
	}
	records, err := utils.SliceAtoi(strings.Fields(lines[1][9:]))
	if err != nil {
		return nil, err
	}
	races := make([]Race, len(times))
	for i := 0; i < len(times); i++ {
		races[i] = Race{
			Time:   times[i],
			Record: records[i],
		}
	}
	return races, nil
}

func PartOne(races []Race) int {
	total := 1
	for _, race := range races {
		min := race.GetMinSpeed()
		max := race.GetMaxSpeed(min)

		total *= max - min + 1
	}
	return total
}

func PartTwo(lines []string) int {
	time, err := strconv.Atoi(strings.Join(strings.Fields(lines[0][7:]), ""))
	if err != nil {
		panic(err)
	}

	distance, err := strconv.Atoi(strings.Join(strings.Fields(lines[1][9:]), ""))
	if err != nil {
		panic(err)
	}
	race := Race{
		Time:   time,
		Record: distance,
	}

	min := race.GetMinSpeed()
	max := race.GetMaxSpeed(min)

	return max - min + 1
}

func Run() {
	lines, err := utils.GetLines("./day_06/input.txt")
	if err != nil {
		panic(err)
	}

	races, err := ParseFile(lines)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Part One: %d\n", PartOne(races))
	fmt.Printf("Part Two: %d\n", PartTwo(lines))
}
