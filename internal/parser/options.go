package parser

import (
	"aoc-2024/internal/utils"
	"errors"
	"flag"
)

type Options struct {
	Input string
	Part  int
	Debug bool
}

func Parse() Options {
	options := Options{
		Input: "",
		Part:  1,
		Debug: false,
	}

	flag.StringVar(&options.Input, "input", "", "relative path to a txt input file")
	flag.IntVar(&options.Part, "part", 1, "part of a aoc day task")
	flag.BoolVar(&options.Debug, "debug", false, "debug mode")
	flag.Parse()

	if len(options.Input) < 1 {
		err := errors.New("path to a txt input file not provided")
		utils.FailOnError(err, "Please provide path to a txt input file")
	}

	return options
}
