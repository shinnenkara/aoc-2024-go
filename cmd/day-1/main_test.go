package main

import (
	"aoc-2024/internal/aoc"
	"aoc-2024/internal/parser"
	"testing"
)

func TestPart1(t *testing.T) {
	options := parser.Options{
		Part:  1,
		Input: "example",
		Debug: false,
	}
	lines := aoc.ParseInput(options)
	result := Part1(lines, options)

	expected := 11
	if result != expected {
		t.Fatalf(`Part 1 example should be: %+v, got: %+v`, expected, result)
	}
}

func TestPart2(t *testing.T) {
	options := parser.Options{
		Part:  2,
		Input: "example",
		Debug: false,
	}
	lines := aoc.ParseInput(options)
	result := Part2(lines, options)

	expected := 31
	if result != expected {
		t.Fatalf(`Part 2 example should be: %+v, got: %+v`, expected, result)
	}
}
