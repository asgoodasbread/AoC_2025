package main

import (
	"adventofcode2025/utils"
	"fmt"
	"math"
	"sort"
)

type Edge struct {
	x int
	y int
}

func main() {
	//part_1()
	part_2()
}

func part_2() {
	input := utils.ReadInputSep("Day9/day9_input.txt", ",")
	edges := parseInput(input)
	//fmt.Println(edges)

	//Calculate compressed coordinates values
	xp := make([]int, 0)
	yp := make([]int, 0)

	for _, edge := range edges {
		xp = append(xp, edge.x)
		yp = append(yp, edge.y)
	}

	xp = getSet(xp)
	yp = getSet(yp)

	sort.Ints(xp)
	sort.Ints(yp)

	//Create a grid
	rows := len(yp)*2 - 1
	cols := len(xp)*2 - 1

	// grid[i][j] == 1 means the block is inside the figure
	grid := make([][]int, rows)
	for i := 0; i < rows; i++ {
		grid[i] = make([]int, cols)
		for j := 0; j < cols; j++ {
			grid[i][j] = 0
		}
	}

	xMap := make(map[int]int)
	for i, v := range xp {
		xMap[v] = i
	}
	yMap := make(map[int]int)
	for i, v := range yp {
		yMap[v] = i
	}

	for i := 0; i < len(edges); i++ {
		// Current Vertex
		p1 := edges[i]
		// Next Vertex (wraps around to 0 at the end)
		p2 := edges[(i+1)%len(edges)]

		// Get Compressed Coordinates from Map
		cx1, cx2 := xMap[p1.x]*2, xMap[p2.x]*2
		cy1, cy2 := yMap[p1.y]*2, yMap[p2.y]*2

		startX, endX := min(cx1, cx2), max(cx1, cx2)
		startY, endY := min(cy1, cy2), max(cy1, cy2)

		for r := startX; r <= endX; r++ {
			for c := startY; c <= endY; c++ {
				grid[r][c] = 1
			}
		}
		// fmt.Printf("Original: (%d,%d)->(%d,%d) \t Compressed: (%d,%d)->(%d,%d)\n",
		// 	p1.x, p1.y, p2.x, p2.y,
		// 	startX, startY, endX, endY)
	}

	// outside := Edge{
	// 	x: -1,
	// 	y: -1,
	// }
	visited := make(map[Edge]bool)

	start := Edge{x: -1, y: -1}
	queue := []Edge{start}
	visited[start] = true
	for len(queue) > 0 {
		// Pop first element
		cur := queue[0]
		queue = queue[1:]

		tx, ty := cur.x, cur.y
		dirs := []Edge{
			{x: tx - 1, y: ty},
			{x: tx + 1, y: ty},
			{x: tx, y: ty - 1},
			{x: tx, y: ty + 1},
		}

		for _, next := range dirs {
			nx, ny := next.x, next.y
			if nx < -1 || ny < -1 || nx > len(grid) || ny > len(grid[0]) {
				continue
			}

			// Hit a wall in the compressed grid
			if nx >= 0 && nx < len(grid) && ny >= 0 && ny < len(grid[0]) {
				if grid[nx][ny] == 1 {
					continue
				}
			}

			if visited[next] {
				continue
			}
			visited[next] = true
			queue = append(queue, next)
		}
	}

	for a := range grid {
		for b := range grid[0] {
			myEdge := Edge{
				x: a,
				y: b,
			}
			if visited[myEdge] == false {
				grid[a][b] = 1
			}
		}
	}

	// for _, r := range grid {
	// 	for _, c := range r {
	// 		fmt.Print(c, " ")
	// 	}
	// 	fmt.Println()
	// }

	//Use PSA (Prefix Sum Array)
	//fmt.Print(strings.Repeat("-", 20), "\n")
	psa := construct_psa(grid)
	// for _, row := range psa {
	// 	for _, x := range row {
	// 		fmt.Print(x, "\t")
	// 	}
	// 	fmt.Println()
	// }

	max := 0
	for i := 0; i < len(edges)-1; i++ {
		for j := i + 1; j < len(edges); j++ {
			p1, p2 := edges[i], edges[j]

			if p1.x == p2.x || p1.y == p2.y {
				continue
			}
			x1 := edges[i].x
			x2 := edges[j].x
			y1 := edges[i].y
			y2 := edges[j].y

			if valid(x1, x2, y1, y2, psa, xMap, yMap) {
				width := math.Abs(float64(x1)-float64(x2)) + 1
				height := math.Abs(float64(y1)-float64(y2)) + 1
				na := width * height
				if int(na) > max {
					max = int(na)
				}
			}
		}
	}
	fmt.Println(max)

}

func valid(x1, x2, y1, y2 int, psa [][]int, xMap, yMap map[int]int) bool {
	cx1, cx2 := xMap[x1]*2, xMap[x2]*2
	cy1, cy2 := yMap[y1]*2, yMap[y2]*2

	startX, endX := min(cx1, cx2), max(cx1, cx2)
	startY, endY := min(cy1, cy2), max(cy1, cy2)

	left := 0
	top := 0
	topleft := 0

	if startX > 0 {
		left = psa[startX-1][endY]
	}
	if startY > 0 {
		top = psa[endX][startY-1]
	}
	if startX > 0 && startY > 0 {
		topleft = psa[startX-1][startY-1]
	}
	count := psa[endX][endY] - left - top + topleft
	return count == (endX-startX+1)*(endY-startY+1)
}

func construct_psa(grid [][]int) [][]int {
	rows := len(grid)
	if rows == 0 {
		return nil
	}
	cols := len(grid[0])
	psa := make([][]int, rows)
	for i := range psa {
		psa[i] = make([]int, cols)
	}
	// 2. Build the PSA
	for x := 0; x < rows; x++ {
		for y := 0; y < cols; y++ {
			left := 0
			top := 0
			topLeft := 0

			if x > 0 {
				left = psa[x-1][y]
			}
			if y > 0 {
				top = psa[x][y-1]
			}
			if x > 0 && y > 0 {
				topLeft = psa[x-1][y-1]
			}
			psa[x][y] = grid[x][y] + left + top - topLeft
		}
	}
	return psa
}

func getSet(input []int) []int {
	myMap := make(map[int]struct{})

	for _, x := range input {
		myMap[x] = struct{}{}
	}
	res := make([]int, 0)
	for x := range myMap {
		res = append(res, x)
	}
	return res
}

func part_1() {
	input := utils.ReadInputSep("Day9/day9_input.txt", ",")
	edges := parseInput(input)
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

func parseInput(input []string) []Edge {
	edges := make([]Edge, 0)

	for i := 0; i < len(input); i += 2 {
		myEdge := Edge{
			x: utils.Atoi(input[i]),
			y: utils.Atoi(input[i+1]),
		}
		edges = append(edges, myEdge)
	}
	return edges
}
