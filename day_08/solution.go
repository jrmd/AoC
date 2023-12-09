package day_08

import (
	"fmt"

	"github.com/jrmd/AoC/utils"
)

type Node struct {
	Left  string
	Right string
}

type Network struct {
	Instr          string
	Nodes          map[string]Node
	StartingPoints []string
}

func (n *Network) PathsToEnd(start string, isKey func(s string) bool) int {
	i := 0
	key := start
	steps := 0

	for {
		instr := n.Instr[i]
		node := n.Nodes[key]

		if instr == 'L' {
			key = node.Left
		} else {
			key = node.Right
		}

		steps++

		if isKey(key) {
			break
		}

		i++
		if i >= len(n.Instr) {
			i = 0
		}
	}

	return steps
}

func ParseFile(lines []string) (*Network, error) {
	network := Network{}

	network.Instr = lines[0]
	network.Nodes = make(map[string]Node)

	for i := 2; i < len(lines); i++ {
		line := lines[i]
		network.Nodes[line[0:3]] = Node{
			Left:  line[7:10],
			Right: line[12:15],
		}

		if line[2] == 'A' {
			network.StartingPoints = append(network.StartingPoints, line[0:3])
		}
	}

	return &network, nil
}

func PartOne(network *Network) int {
	return network.PathsToEnd("AAA", func(s string) bool {
		return s == "ZZZ"
	})
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(integers ...int) int {
	a := integers[0]
	b := integers[1]
	result := a * b / GCD(a, b)

	for i := 2; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func PartTwo(network *Network) int {
	currentKeys := network.StartingPoints
	total := len(currentKeys)

	multiples := make([]int, total)

	for i, start := range network.StartingPoints {
		multiples[i] = network.PathsToEnd(start, func(s string) bool {
			return s[2] == 'Z'
		})

	}

	return LCM(multiples...)
}

func Run() {
	lines, err := utils.GetLines("./day_08/input.txt")
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
