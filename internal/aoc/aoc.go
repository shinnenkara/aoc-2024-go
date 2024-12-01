package aoc

import (
	"aoc-2024/internal/parser"
	"aoc-2024/internal/utils"
	"fmt"
)

type Day struct {
	Part1 func(lines []string, options parser.Options) int
	Part2 func(lines []string, options parser.Options) int
}

func (day *Day) Solve() {
	options := parser.Parse()
	fmt.Printf("%+v\n", options)

	lines := ParseInput(options)
	utils.PrintOnDebug(options.Debug, fmt.Sprintf("%+v", lines))

	var result int
	if options.Part == 1 {
		result = day.Part1(lines, options)
	} else {
		result = day.Part2(lines, options)
	}

	fmt.Printf("Result for the part %+v is: %+v\n", options.Part, result)
}

func ParseInput(options parser.Options) []string {
	lines, err := parser.ParseTxt(options.Input)
	utils.FailOnError(err, "cannot parse provided txt input file")

	return lines
}
