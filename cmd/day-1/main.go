package main

import (
	"aoc-2024/internal/aoc"
	"aoc-2024/internal/parser"
	"aoc-2024/internal/utils"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

func main() {
	day1 := aoc.Day{Part1: Part1, Part2: Part2}
	day1.Solve()
}

func Part1(lines []string, options parser.Options) int {
	list1, list2 := prepareLists(lines, options)
	if len(list1) != len(list2) {
		panic("Lengths not equal")
	}
	length := len(list1)

	sort.Ints(list1)
	utils.PrintOnDebug(options.Debug, fmt.Sprintf("sorted list1: %+v\n", list1))
	sort.Ints(list2)
	utils.PrintOnDebug(options.Debug, fmt.Sprintf("sorted list2: %+v\n", list2))

	var distances []int
	for i := 0; i < length; i++ {
		val1 := list1[i]
		val2 := list2[i]
		distance := int(math.Abs(float64(val1 - val2)))
		distances = append(distances, distance)
	}
	utils.PrintOnDebug(options.Debug, fmt.Sprintf("distances: %+v\n", distances))

	return utils.Sum(distances)
}

func Part2(lines []string, options parser.Options) int {
	list1, list2 := prepareLists(lines, options)

	var similarityScores []int
	for i := 0; i < len(list1); i++ {
		val1 := list1[i]
		similarityMult := 0
		for j := 0; j < len(list2); j++ {
			val2 := list2[j]
			if val1 == val2 {
				similarityMult++
			}
		}

		similarityScore := val1 * similarityMult
		similarityScores = append(similarityScores, similarityScore)
	}
	utils.PrintOnDebug(options.Debug, fmt.Sprintf("similarity scores: %+v\n", similarityScores))

	return utils.Sum(similarityScores)
}

func prepareLists(lines []string, options parser.Options) (list1 []int, list2 []int) {
	for _, l := range lines {
		chars := strings.Split(l, "   ")
		utils.PrintOnDebug(options.Debug, fmt.Sprintf("chars: %v\n", chars))

		num1, _ := strconv.Atoi(chars[0])
		list1 = append(list1, num1)

		num2, _ := strconv.Atoi(chars[1])
		list2 = append(list2, num2)
	}
	utils.PrintOnDebug(options.Debug, fmt.Sprintf("list1: %+v\n", list1))
	utils.PrintOnDebug(options.Debug, fmt.Sprintf("list2: %+v\n", list2))

	return list1, list2
}
