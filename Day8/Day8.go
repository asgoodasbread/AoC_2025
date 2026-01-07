package main

import (
	"adventofcode2025/utils"
	"fmt"
	"math"
	"sort"
)

type Jbox struct {
	x int
	y int
	z int
}

func main() {
	part_1()
	fmt.Println()
}

func part_1() {
	input := utils.ReadInputSep("Day8/day8_input.txt", ",")
	var jboxes []Jbox

	//Parse input into structs
	for i := 0; i < len(input); i += 3 {
		myJbox := Jbox{
			x: utils.Atoi(input[i]),
			y: utils.Atoi(input[i+1]),
			z: utils.Atoi(input[i+2]),
		}
		jboxes = append(jboxes, myJbox)
	}

	boxDist := make(map[[2]int]float64, 0)
	for i := 0; i < len(jboxes)-1; i++ {
		for j := i + 1; j < len(jboxes); j++ {
			boxDist[[2]int{i, j}] = calculateDistance(jboxes[i], jboxes[j])
		}
	}

	sortedEdges := make([][2]int, 0, len(boxDist))

	for edge := range boxDist {
		sortedEdges = append(sortedEdges, [2]int{edge[0], edge[1]})
	}

	sort.Slice(sortedEdges, func(i, j int) bool {
		return boxDist[sortedEdges[i]] < boxDist[sortedEdges[j]]
	})

	//Create parents and initialize
	parents := make([]int, len(jboxes))
	for i := range len(jboxes) {
		parents[i] = i
	}

	//1000 shortest distances
	for _, edge := range sortedEdges[:1000] {
		union(edge[0], edge[1], parents)
	}

	groupSizes := make(map[int]int)

	for i := 0; i < len(parents); i++ {
		root := findRoot(i, parents)
		groupSizes[root]++
	}

	sortedGroupSizes := make([]int, len(groupSizes))
	for _, size := range groupSizes {
		sortedGroupSizes = append(sortedGroupSizes, size)
	}

	sort.Slice(sortedGroupSizes, func(i, j int) bool {
		return sortedGroupSizes[i] > sortedGroupSizes[j]
	})

	fmt.Println(sortedGroupSizes[0] * sortedGroupSizes[1] * sortedGroupSizes[2])

}

func findRoot(i int, parents []int) int {
	if parents[i] == i {
		return i
	}
	return findRoot(parents[i], parents)
}

func union(i, j int, parents []int) {
	rootI := findRoot(i, parents)
	rootJ := findRoot(j, parents)

	if rootI != rootJ {
		parents[rootI] = rootJ
	}
}

func calculateDistance(sJbox Jbox, eJbox Jbox) float64 {
	xdif := float64(sJbox.x - eJbox.x)
	ydif := float64(sJbox.y - eJbox.y)
	zdif := float64(sJbox.z - eJbox.z)

	return math.Sqrt(math.Pow(xdif, 2.0) + math.Pow(ydif, 2.0) + math.Pow(zdif, 2.0))
}
