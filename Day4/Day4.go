package main

import (
	"adventofcode2025/utils"
	"fmt"
)

type Pos struct {
	x int
	y int
}

func main() {
	//fmt.Println("Day 4 !")
	//part_1()
	part_2()
}

func part_2() {
	input_grid := utils.ReadAsGrid("Day4/day4_input.txt")
	fmt.Println(reduce_grid(input_grid))
}

func reduce_grid(grid [][]rune) int {
	removed := 0
	for x, line := range grid {
		for y, elem := range line {
			curPos := Pos{
				x: x,
				y: y,
			}
			if elem == '@' {
				if check_adjacent(curPos, grid) {
					removed++
					grid[curPos.x][curPos.y] = '.'
				}
			}
		}
	}
	if removed == 0 {
		return 0
	}
	return removed + reduce_grid(grid)
}

func part_1() {
	input_grid := utils.ReadAsGrid("Day4/day4_input.txt")
	res := 0
	for x, line := range input_grid {
		for y, elem := range line {
			curPos := Pos{
				x: x,
				y: y,
			}
			if elem == '@' {
				if check_adjacent(curPos, input_grid) {
					res++
				}
			}
		}
	}
	fmt.Println(res)
}

func check_adjacent(curPos Pos, grid [][]rune) bool {
	toilet_paper := 0
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			if i == 0 && j == 0 {
				continue
			}
			nx, ny := curPos.x+i, curPos.y+j
			if nx < len(grid[0]) && nx >= 0 && ny < len(grid) && ny >= 0 && grid[nx][ny] == rune('@') {
				toilet_paper++
			}
			if toilet_paper >= 4 {
				return false
			}
		}
	}
	//fmt.Println(curPos.x, curPos.y)
	return true
}
