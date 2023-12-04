package day_02

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jrmd/AoC/utils"
)

type Set struct {
	Red   int
	Blue  int
	Green int
}

type Game struct {
	Id   int
	Sets []*Set
}

func (game *Game) Valid(r int, g int, b int) bool {
	for _, set := range game.Sets {
		if set.Red > r || set.Green > g || set.Blue > b {
			return false
		}
	}

	return true
}

func (game *Game) Largest() (int, int, int) {
	red, green, blue := 0, 0, 0
	for _, set := range game.Sets {
		if set.Blue > blue {
			blue = set.Blue
		}
		if set.Red > red {
			red = set.Red
		}
		if set.Green > green {
			green = set.Green
		}
	}

	return red, green, blue
}

func (game *Game) Power() int {
	r, g, b := game.Largest()

	return r * g * b
}

func ParseSet(set string) *Set {
	cubes := strings.Split(set, ", ")
	red, blue, green := 0, 0, 0
	for _, cube := range cubes {
		splitCube := strings.Split(cube, " ")
		count, colour := splitCube[0], splitCube[1]

		total, err := strconv.Atoi(count)
		if err != nil {
			continue
		}

		switch colour {
		case "red":
			red = total
		case "green":
			green = total
		case "blue":
			blue = total
		}
	}

	return &Set{
		red,
		blue,
		green,
	}
}

func ParseLine(input string) (*Game, error) {
	// Game 1: 1 red, 2 blue, 3 green; 3 red, 2 blue, 1 green;
	result := strings.Split(input, ": ")
	sid, game := result[0], result[1]

	id, err := strconv.Atoi(sid[5:])
	if err != nil {
		return nil, err
	}

	sets := strings.Split(game, "; ")

	parsedSets := []*Set{}

	for _, set := range sets {
		parsedSets = append(parsedSets, ParseSet(set))
	}

	return &Game{
		Id:   id,
		Sets: parsedSets,
	}, nil
}

func ParseFile(input []string) ([]*Game, error) {
	games := []*Game{}
	for _, value := range input {
		game, err := ParseLine(value)
		if err != nil {
			return nil, err
		}
		games = append(games, game)
	}

	return games, nil
}

func PartOne(games []*Game) int {
	const RED = 12
	const BLUE = 14
	const GREEN = 13

	total := 0

	for _, game := range games {
		if game.Valid(RED, GREEN, BLUE) {
			total += game.Id
			continue
		}
	}

	return total
}

func PartTwo(games []*Game) int {
	total := 0

	for _, game := range games {
		total += game.Power()
	}

	return total
}

func Run() {
	lines, err := utils.GetLines("./day_02/input.txt")
	if err != nil {
		panic(err)
	}

	games, err := ParseFile(lines)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Part One: %d\n", PartOne(games))
	fmt.Printf("Part Two: %d\n", PartTwo(games))
}
