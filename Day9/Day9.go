package main

import (
	"adventofcode2025/utils"
	"fmt"
	"math"
)

type Edge struct {
	x int
	y int
}

func main() {
	part_1()
	part_2()
}

func part_2() {
	input := utils.ReadInputSep("Day9/day9_input.txt", ",")
	fmt.Println(input)
}

func part_1() {
	input := utils.ReadInputSep("Day9/day9_input.txt", ",")
	edges := make([]Edge, 0)

	for i := 0; i < len(input); i += 2 {
		myEdge := Edge{
			x: utils.Atoi(input[i]),
			y: utils.Atoi(input[i+1]),
		}
		edges = append(edges, myEdge)
	}
	//fmt.Println(edges)

	area := float64(0)
	for i := 0; i < len(edges)-1; i++ {
		for j := i + 1; j < len(edges); j++ {
			na := math.Abs(float64(edges[i].x)-float64(edges[j].x)+1) * math.Abs(float64(edges[i].y)-float64(edges[j].y)+1)
			if na >= area {
				area = na
				//fmt.Println(edges[i].x, edges[j].x, edges[i].y, edges[j].y)
			}
		}
	}
	fmt.Println(int(area))
}
