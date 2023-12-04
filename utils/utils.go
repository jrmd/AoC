package utils

import (
	"os"
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
