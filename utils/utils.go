package utils

import (
	"os"
	"strconv"
	"strings"
)

func GetLines(file string) ([]string, error) {
	inputBytes, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	input := string(inputBytes)
	input = strings.TrimSpace(input)
	lines := strings.Split(input, "\n")

	return lines, nil
}

func SliceAtoi(sa []string) ([]int, error) {
	si := make([]int, 0, len(sa))
	for _, a := range sa {
		i, err := strconv.Atoi(a)
		if err != nil {
			return si, err
		}
		si = append(si, i)
	}
	return si, nil
}
