package main

import (
	"aoc-2024/internal/aoc"
	"aoc-2024/internal/parser"
	"aoc-2024/internal/utils"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	day2 := aoc.Day{Part1: Part1, Part2: Part2}
	day2.Solve()
}

func Part1(lines []string, options parser.Options) int {
	lists := prepareLists(lines, options)
	utils.PrintOnDebug(options.Debug, fmt.Sprintf("lines parsed: %+v\n", len(lists)))

	verifiedLists := verifyLists(lists, 0, options)
	utils.PrintOnDebug(options.Debug, fmt.Sprintf("verified: %+v\n", len(verifiedLists)))

	return len(verifiedLists)
}

func Part2(lines []string, options parser.Options) int {
	lists := prepareLists(lines, options)
	utils.PrintOnDebug(options.Debug, fmt.Sprintf("lines parsed: %+v\n", len(lists)))

	verifiedLists := verifyLists(lists, 1, options)
	utils.PrintOnDebug(options.Debug, fmt.Sprintf("verified: %+v\n", len(verifiedLists)))

	return len(verifiedLists)
}

func prepareLists(lines []string, options parser.Options) (lists [][]int) {
	for _, l := range lines {
		chars := strings.Split(l, " ")
		utils.PrintOnDebug(options.Debug, fmt.Sprintf("chars: %v\n", chars))

		var list []int
		for _, char := range chars {
			num, _ := strconv.Atoi(char)
			list = append(list, num)
		}
		lists = append(lists, list)
		utils.PrintOnDebug(options.Debug, fmt.Sprintf("list: %+v\n", list))
	}

	return lists
}

func verifyLists(lists [][]int, tolerate int, options parser.Options) (verifiedLists [][]int) {
	for _, list := range lists {
		safe := isSafe(list, options)
		if !safe {
			safe = trySafe(list, 0, tolerate, options)
		}

		if safe {
			verifiedLists = append(verifiedLists, list)
		}
	}

	return verifiedLists
}

func trySafe(list []int, count int, tolerate int, options parser.Options) bool {
	if count >= tolerate {
		return false
	}

	for i := 0; i < len(list); i++ {
		partialList := preparePartialList(list, i, options)
		safe := isSafe(partialList, options)
		if !safe {
			safe = trySafe(partialList, count+1, tolerate, options)
		}

		if safe {
			return true
		}
	}

	return false
}

func isSafe(list []int, options parser.Options) bool {
	increasing := isIncreasing(list, options)
	if increasing {
		return isSafelyChanged(list, options)
	}

	decreasing := isDecreasing(list, options)
	if decreasing {
		return isSafelyChanged(list, options)
	}
	utils.PrintOnDebug(options.Debug, fmt.Sprintf("list is not safe: %+v\n", list))

	return false
}

func preparePartialList(list []int, index int, options parser.Options) (partialList []int) {
	if index >= len(list) {
		panic("index out of range")
	}

	var part1 []int
	if index > 0 {
		part1 = list[:index]
	}
	partialList = append(partialList, part1...)

	var part2 []int
	if index < len(list)-1 {
		part2 = list[index+1:]
	}
	partialList = append(partialList, part2...)

	return partialList
}

func isIncreasing(list []int, options parser.Options) bool {
	for i := 0; i < len(list)-1; i++ {
		a := list[i]
		b := list[i+1]
		//utils.PrintOnDebug(options.Debug, fmt.Sprintf("a, b: %+v %+v\n", a, b))
		if a >= b {
			return false
		}
	}
	utils.PrintOnDebug(options.Debug, fmt.Sprintf("list increasing: %+v\n", list))

	return true
}

func isDecreasing(list []int, options parser.Options) bool {
	for i := 0; i < len(list)-1; i++ {
		a := list[i]
		b := list[i+1]
		if a <= b {
			return false
		}
	}
	utils.PrintOnDebug(options.Debug, fmt.Sprintf("list decreasing: %+v\n", list))

	return true
}

func isSafelyChanged(list []int, options parser.Options) bool {
	for i := 0; i < len(list)-1; i++ {
		a := list[i]
		b := list[i+1]
		if !isGraduallyChanged(a, b, options) {
			utils.PrintOnDebug(options.Debug, fmt.Sprintf("list is not safely changed: %+v\n", list))
			return false
		}
	}
	utils.PrintOnDebug(options.Debug, fmt.Sprintf("list safely changed: %+v\n", list))

	return true
}

func isGraduallyChanged(a int, b int, options parser.Options) bool {
	diff := math.Abs(float64(a - b))
	utils.PrintOnDebug(options.Debug, fmt.Sprintf("diff is: %+v\n", diff))

	return diff <= 3
}
