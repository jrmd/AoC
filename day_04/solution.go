package day_04

import (
	"fmt"
	"math"
	"strings"

	"github.com/jrmd/AoC/utils"
)

type Game struct {
	Index   int
	Winners []string
	Numbers []string
}

func (g *Game) Wins() int {
	count := 0
	for _, win := range g.Winners {
		for _, number := range g.Numbers {
			if number == win {
				count += 1
				break
			}
		}
	}

	return count
}

func (g *Game) Score() int {
	count := g.Wins()

	if count == 0 {
		return count
	}

	return int(math.Pow(2, float64(count-1)))
}

func ParseFile(lines []string) ([]Game, error) {
	games := []Game{}

	for i, line := range lines {
		pieces := strings.Split(line, ": ")
		game := strings.Split(pieces[1], " | ")
		games = append(games, Game{
			Index:   i,
			Winners: strings.Fields(game[0]),
			Numbers: strings.Fields(game[1]),
		})
	}

	return games, nil
}

func PartOne(games []Game) int {
	total := 0

	for _, game := range games {
		total += game.Score()
	}

	return total
}

func PartTwo(games []Game) int {
	total := 0

	cards := make([]int, len(games))
	for i := range cards {
		cards[i] = 1
	}

	for i, game := range games {
		currWins := game.Wins()
		totalCards := cards[i]

		for j := 1; j <= currWins && j < len(cards); j++ {
			cards[i+j] += totalCards
		}
	}

	for _, v := range cards {
		total += v
	}

	return total
}

func Run() {
	lines, err := utils.GetLines("./day_04/input.txt")
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
