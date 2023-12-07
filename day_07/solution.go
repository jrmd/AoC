package day_07

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/jrmd/AoC/utils"
)

const (
	FIVE_KIND  = 0
	FOUR_KIND  = 1
	FULL_HOUSE = 2
	THREE_KIND = 3
	TWO_PAIR   = 4
	ONE_PAIR   = 5
	HIGH_CARD  = 6
)

type Game struct {
	Cards  string
	Points int
	Type   int
}

func GetCardType(cards string) int {
	cardCount := map[string]int{}

	for _, card := range cards {
		cardCount[string(card)]++
	}

	if len(cardCount) == 1 {
		return FIVE_KIND
	}

	if len(cardCount) == 2 {
		for _, count := range cardCount {
			if count == 4 {
				return FOUR_KIND
			}
			if count == 3 {
				return FULL_HOUSE
			}
		}
	}

	if len(cardCount) == 3 {
		for _, count := range cardCount {
			if count == 3 {
				return THREE_KIND
			}
			if count == 2 {
				return TWO_PAIR
			}
		}
	}

	if len(cardCount) == 4 {
		return ONE_PAIR
	}

	return HIGH_CARD
}

func GetJokerType(cards string) int {
	cardCount := map[string]int{}
	jokerCount := 0

	for _, card := range cards {
		if string(card) == "J" {
			jokerCount++
			continue
		}
		cardCount[string(card)]++
	}

	if jokerCount > 0 {
		// add joker to highest cardCount
		highest := 0
		highestIndex := "J"
		for i, count := range cardCount {
			if count > highest {
				highest = count
				highestIndex = i
			}
		}

		cardCount[highestIndex] += jokerCount
	}

	if len(cardCount) == 1 {
		return FIVE_KIND
	}

	if len(cardCount) == 2 {
		for _, count := range cardCount {
			if count == 4 {
				return FOUR_KIND
			}
			if count == 3 {
				return FULL_HOUSE
			}
		}
	}

	if len(cardCount) == 3 {
		for _, count := range cardCount {
			if count == 3 {
				return THREE_KIND
			}
			if count == 2 {
				return TWO_PAIR
			}
		}
	}

	if len(cardCount) == 4 {
		return ONE_PAIR
	}

	return HIGH_CARD
}

func ParseFile(lines []string, withJokers bool) ([]Game, error) {
	games := make([]Game, len(lines))

	for i, line := range lines {
		pieces := strings.Fields(line)
		points, err := strconv.Atoi(pieces[1])
		if err != nil {
			return nil, err
		}

		var cardType int
		if withJokers {
			cardType = GetJokerType(pieces[0])
		} else {
			cardType = GetCardType(pieces[0])
		}

		games[i] = Game{
			Cards:  pieces[0],
			Points: points,
			Type:   cardType,
		}
	}

	return games, nil
}

func SortGames(games []Game, CARD_MAP map[byte]int) {
	slices.SortFunc(games, func(a, b Game) int {
		if CARD_MAP[a.Cards[0]] > CARD_MAP[b.Cards[0]] {
			return -1
		}

		if CARD_MAP[a.Cards[0]] < CARD_MAP[b.Cards[0]] {
			return 1
		}

		for i := 1; i < len(a.Cards); i++ {
			if CARD_MAP[a.Cards[i]] > CARD_MAP[b.Cards[i]] {
				return -1
			}

			if CARD_MAP[a.Cards[i]] < CARD_MAP[b.Cards[i]] {
				return 1
			}
		}

		return 0
	})
}

func PartOne(lines []string) int {
	games, err := ParseFile(lines, false)
	if err != nil {
		panic(err)
	}

	var (
		CARD_MAP = map[byte]int{
			'2': 0,
			'3': 1,
			'4': 2,
			'5': 3,
			'6': 4,
			'7': 5,
			'8': 6,
			'9': 7,
			'T': 8,
			'J': 9,
			'Q': 10,
			'K': 11,
			'A': 12,
		}
	)

	piles := make([][]Game, 7)

	for _, game := range games {
		piles[game.Type] = append(piles[game.Type], game)
	}

	total := 0

	for i := 0; i < len(piles); i++ {
		if len(piles[i]) > 0 {
			SortGames(piles[i], CARD_MAP)
		}
	}

	rank := 0

	for i := len(piles) - 1; i >= 0; i-- {
		if len(piles[i]) > 0 {
			for j := len(piles[i]) - 1; j >= 0; j-- {
				total += piles[i][j].Points * (rank + 1)
				rank++
			}
		}
	}

	return total
}

func PartTwo(lines []string) int {
	games, err := ParseFile(lines, true)
	if err != nil {
		panic(err)
	}

	var (
		CARD_MAP = map[byte]int{
			'J': 0,
			'2': 1,
			'3': 2,
			'4': 3,
			'5': 4,
			'6': 5,
			'7': 6,
			'8': 7,
			'9': 8,
			'T': 9,
			'Q': 10,
			'K': 11,
			'A': 12,
		}
	)

	piles := make([][]Game, 7)

	for _, game := range games {
		piles[game.Type] = append(piles[game.Type], game)
	}

	total := 0

	for i := 0; i < len(piles); i++ {
		if len(piles[i]) > 0 {
			SortGames(piles[i], CARD_MAP)
		}
	}

	rank := 0

	for i := len(piles) - 1; i >= 0; i-- {
		if len(piles[i]) > 0 {
			for j := len(piles[i]) - 1; j >= 0; j-- {
				total += piles[i][j].Points * (rank + 1)
				rank++
			}
		}
	}

	return total
}

func Run() {
	lines, err := utils.GetLines("./day_07/input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Part One: %d\n", PartOne(lines))
	fmt.Printf("Part Two: %d\n", PartTwo(lines))
}
