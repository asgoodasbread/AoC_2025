package utils

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func ReadDay10(file_name string) [][]string {
	var res [][]string
	str, err := os.ReadFile(file_name)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	content := string(str)
	content = strings.ReplaceAll(content, "\r", "")

	//fmt.Println(string(content))
	lines := strings.Split(string(content), "\n")

	for _, line := range lines {
		var newLines []string
		lightsLastIndex := 0
		joltageStartIndex := 0
		for i, char := range line {
			switch char {
			case ']':
				lightsLastIndex = i
			case '{':
				joltageStartIndex = i
			}
		}
		newLines = append(newLines, line[1:lightsLastIndex], line[lightsLastIndex+1:joltageStartIndex], line[joltageStartIndex:])
		res = append(res, newLines)
	}
	return res
}

func ReadInputSep(file_name string, separator string) []string {
	content, err := os.ReadFile(file_name)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	str := string(content)

	str = strings.ReplaceAll(str, "\n", separator)
	str = strings.ReplaceAll(str, "\r", "")

	parts := strings.Split(str, separator)
	return parts
}

func Atoi(s string) int {
	val, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("Error converting '%s' to int: %v", s, err)
	}
	return val
}

func ReadAsGrid(file_name string) [][]rune {
	var grid [][]rune
	file, err := os.Open(file_name)
	if err != nil {
		log.Panicln("Error opening file !")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		var grid_line []rune
		for _, c := range line {
			grid_line = append(grid_line, rune(c))
		}
		grid = append(grid, grid_line)
	}
	return grid
}

func ReadDay5(file_name string) ([]string, []int) {
	id_ranges := make([]string, 0)
	ids := make([]int, 0)

	file, err := os.Open(file_name)
	if err != nil {
		log.Panicln("Error opening file !")
	}
	scanner := bufio.NewScanner(file)
	flag := false

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			flag = true
			continue
		}

		if !flag {
			id_ranges = append(id_ranges, line)
		} else {
			ids = append(ids, Atoi(line))
		}
	}
	return id_ranges, ids
}

// func main() {
// 	input := ReadInputSep("Day2/test_day2.txt", ",")
// 	//fmt.Println([]rune(input[5]))
// 	for i, x := range input {
// 		if i == 7 {
// 			fmt.Println()
// 		}
// 		//fmt.Println(x)
// 	}
// }
