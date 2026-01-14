package main

import (
	"adventofcode2025/utils"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	part_1()
}

func part_1() {
	input := utils.ReadInputSep("Day12/day12_input.txt", "\n\n")
	splitIndex := findSplitIndex(input)
	input = input[splitIndex:]
	res := 0
	re := regexp.MustCompile(`\d+`)
	for _, line := range input {
		matches := re.FindAllString(line, -1)
		iMatch := toInt(matches)
		x, y, presents := iMatch[0], iMatch[1], iMatch[2:]
		pSum := 0
		for _, a := range presents {
			pSum += a
		}

		if (math.Floor(float64(x/3)))*math.Floor(float64(y/3)) >= float64(pSum) {
			res += 1
		}

	}
	fmt.Println(res)
}

func toInt(x []string) []int {
	res := make([]int, 0, len(x))
	for _, el := range x {
		val, _ := strconv.Atoi(el)
		res = append(res, val)
	}
	return res
}

func findSplitIndex(input []string) int {
	for i, line := range input {
		if strings.Contains(line, "x") {
			return i
		}
	}
	return -1
}
