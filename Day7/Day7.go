package main

import (
	"adventofcode2025/utils"
	"fmt"
)

func main() {
	part_1()
	part_2()
}

func part_2() {
	grid := utils.ReadAsGrid("Day7/day7_input.txt")
	memo := make(map[[2]int]int)
	//Find 'S'
	startX, startY := 0, 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 'S' {
				startX, startY = i, j
			}
		}
	}
	fmt.Println(getTimelines(startX, startY, memo, grid))
}

func getTimelines(r int, c int, memo map[[2]int]int, grid [][]rune) int {
	key := [2]int{r, c}

	if val, ok := memo[key]; ok {
		return val
	}
	if r >= len(grid) {
		return 1
	}
	var result int

	currentChar := grid[r][c]
	switch currentChar {
	case '.', 'S':
		result = getTimelines(r+1, c, memo, grid)
	case '^':
		result = getTimelines(r, c-1, memo, grid) + getTimelines(r, c+1, memo, grid)
	}

	memo[key] = result
	return result
}

func part_1() {
	input := utils.ReadAsGrid("Day7/day7_input.txt")
	visited := make([][]bool, len(input))
	for i := range visited {
		visited[i] = make([]bool, len(input[0]))
	}
	//Find 'S' and update grid
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[0]); j++ {
			if string(input[i][j]) == "S" {
				updateGrid(input, i+1, j, visited)
				goto CheckFinalGrid
			}
		}

	}
CheckFinalGrid:
	//Print final grid
	// for _, line := range input {
	// 	fmt.Println(string(line))
	// }

	res := 0
	//Count the split beams
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[0]); j++ {
			if input[i][j] == '^' && input[i-1][j] == '|' {
				res++
			}
		}
	}
	fmt.Println(res)
}

func updateGrid(grid [][]rune, x int, y int, visited [][]bool) {
	if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[0]) {
		return
	}

	if visited[x][y] {
		return
	}
	visited[x][y] = true

	//Handle start
	switch grid[x][y] {
	case '.':
		grid[x][y] = '|'
		updateGrid(grid, x+1, y, visited)
	case '|':
		updateGrid(grid, x+1, y, visited)
	//Handle split
	case '^':
		updateGrid(grid, x, y-1, visited)
		updateGrid(grid, x, y+1, visited)
	default:
		return
	}
	return
}
