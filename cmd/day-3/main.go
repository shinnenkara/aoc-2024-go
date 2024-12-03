package main

import (
	"aoc-2024/internal/aoc"
	"aoc-2024/internal/parser"
	"aoc-2024/internal/utils"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	day3 := aoc.Day{Part1: Part1, Part2: Part2}
	day3.Solve()
}

func Part1(lines []string, options parser.Options) int {
	line := prepareLine(lines, options)
	result := parseForMul(line, options)

	return result
}

func Part2(lines []string, options parser.Options) int {
	line := prepareLine(lines, options)
	result := parse(line, options)

	return result
}

func prepareLine(lines []string, options parser.Options) string {
	var line string
	for _, l := range lines {
		line += l
	}
	utils.PrintOnDebug(options.Debug, fmt.Sprintf("line: %+v\n", line))

	return line
}

func parseForMul(line string, options parser.Options) (sum int) {
	reg := regexp.MustCompile("mul\\((\\d{1,3}),(\\d{1,3})\\)")

	matches := reg.FindAllString(line, -1)
	for _, m := range matches {
		utils.PrintOnDebug(options.Debug, fmt.Sprintf("match: %+v\n", m))
		nums := strings.Split(m, ",")
		str1, _ := strings.CutPrefix(nums[0], "mul(")
		str2, _ := strings.CutSuffix(nums[1], ")")
		utils.PrintOnDebug(options.Debug, fmt.Sprintf("num 1: %+v\n", str1))
		utils.PrintOnDebug(options.Debug, fmt.Sprintf("num 2: %+v\n", str2))

		num1, _ := strconv.Atoi(str1)
		num2, _ := strconv.Atoi(str2)
		sum += num1 * num2
	}

	return sum
}

func parse(line string, options parser.Options) (sum int) {
	doBlocks := strings.Split(line, "do()")

	for _, block := range doBlocks {
		correctBlocks := strings.Split(block, "don't()")
		result := parseForMul(correctBlocks[0], options)

		sum += result
	}

	return sum
}
