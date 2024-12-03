package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/jrmd/AoC/day_01"

	"github.com/urfave/cli/v2"
)

func CreateDay(ctx *cli.Context) error {
	// https://pkg.go.dev/github.com/otiai10/copy#section-readme
	return nil
}

func RunDay(ctx *cli.Context) error {
	arg := ctx.Args().First()
	runAll := ctx.Bool("all")

	days := []func(){
		day_01.Run,
  }

	if runAll {
		for i, solution := range days {
			fmt.Printf("--- Day %d ---\n", i+1)
			solution()
		}
		return nil
	}

	argParsed := strings.Split(arg, ",")

	for _, i := range argParsed {
		index, err := strconv.Atoi(i)

		if err != nil || index < 1 || index > len(days) {
			return errors.New("invalid day to run")
		}

		fmt.Printf("--- Day %d ---\n", index)
		days[index-1]()
	}

	return nil
}

func main() {
	app := &cli.App{
		Name:  "aoc",
		Usage: "Handy utility to run AoC components",
		Commands: []*cli.Command{
			{
				Name:    "run",
				Aliases: []string{"r"},
				Usage:   "Run a AoC task",
				Action:  RunDay,
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:    "all",
						Aliases: []string{"A"},
						Usage:   "Run all days",
					},
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
