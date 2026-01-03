package main

import (
	"adventofcode2025/utils"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	//part_1()
	part_2()
}

func part_2() {
	grid := utils.ReadAsGrid("Day6/day6_input.txt")
	ops := grid[len(grid)-1]
	numberGrid := grid[:len(grid)-1]

	cols := len(numberGrid[0])
	rows := len(numberGrid)

	res := uint64(0)

	// Lambda to handle each block
	var currentBlockNums []uint64
	currentBlockOp := ' '
	processBlock := func() {
		if len(currentBlockNums) == 0 {
			return
		}
		var blockResult uint64
		if currentBlockOp == '+' {
			blockResult = 0
			for _, n := range currentBlockNums {
				blockResult += n
			}
		} else {
			blockResult = 1
			for _, n := range currentBlockNums {
				blockResult *= n
			}
		}
		res += blockResult
		currentBlockNums = []uint64{}
		currentBlockOp = ' '
	}

	for x := 0; x < cols; x++ {
		var sb strings.Builder
		isEmptyColumn := true

		for y := 0; y < rows; y++ {
			if x < len(numberGrid[y]) {
				r := rune(numberGrid[y][x])
				if !unicode.IsSpace(r) {
					sb.WriteRune(r)
					isEmptyColumn = false
				}
			}
		}

		if isEmptyColumn {
			processBlock()
		} else {
			val, _ := strconv.Atoi(sb.String())
			currentBlockNums = append(currentBlockNums, uint64(val))

			if x < len(ops) {
				opChar := rune(ops[x])
				if opChar == '+' || opChar == '*' {
					currentBlockOp = opChar
				}
			}
		}
	}
	processBlock()
	fmt.Println(res)
}

func part_1() {
	lines := getLines()
	input := utils.ReadInputSep("Day6/day6_input.txt", " ")
	clean_input := cleanInput(input)

	operations := clean_input[(len(clean_input) - len(clean_input)/len(lines)):]
	var res uint64 = 0
	for i := 0; i < len(clean_input)/len(lines); i++ {
		if operations[i] == "*" {
			x := uint64(1)
			for j := 0; j < len(clean_input)/(len(clean_input)/len(lines))-1; j++ {
				val := uint64(utils.Atoi(clean_input[(j*(len(clean_input)/len(lines)))+i]))
				x *= val
			}
			res += x
		} else {
			x := uint64(0)
			for j := 0; j < len(clean_input)/(len(clean_input)/len(lines))-1; j++ {
				val := uint64(utils.Atoi(clean_input[(j*(len(clean_input)/len(lines)))+i]))
				x += val
			}
			res += x
		}

	}
	fmt.Println(res)
}

func getLines() []string {
	//Get number of lines (rows)
	contentBytes, _ := os.ReadFile("Day6/test_input.txt")
	content := string(contentBytes)
	lines := strings.Split(content, "\n")
	return lines
}

func cleanInput(input []string) []string {
	var clean_input []string

	//Clean the input
	for _, x := range input {
		cleaned := strings.ReplaceAll(x, " ", "")

		if cleaned != "" {
			clean_input = append(clean_input, cleaned)
		}
	}
	return clean_input
}
