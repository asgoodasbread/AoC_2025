package main

import (
	"adventofcode2025/utils"
	"fmt"
	"strings"
)

func main() {
	//part_1()
	part_2()
}

func part_2() {
	input := utils.ReadInputSep("Day11/day11_input.txt", ":")
	devices := parseDevices(input)
	visited := make(map[string]bool)
	fmt.Println(countUniquePaths("svr", "fft", devices, visited)*
		countUniquePaths("fft", "dac", devices, visited)*
		countUniquePaths("dac", "out", devices, visited) +
		countUniquePaths("svr", "dac", devices, visited)*
			countUniquePaths("dac", "fft", devices, visited)*
			countUniquePaths("fft", "out", devices, visited))
}

var cache = make(map[string]int)

func countUniquePaths(source string, dest string, devices map[string][]string, visited map[string]bool) int {
	if source == dest {
		return 1
	}
	cacheKey := source + "->" + dest
	if val, exists := cache[cacheKey]; exists {
		return val
	}
	if visited[source] {
		return 0
	}

	visited[source] = true
	totalPaths := 0

	for _, x := range devices[source] {
		totalPaths += countUniquePaths(x, dest, devices, visited)
	}
	visited[source] = false
	cache[cacheKey] = totalPaths
	return totalPaths
}

func part_1() {
	input := utils.ReadInputSep("Day11/day11_input.txt", ":")
	devices := parseDevices(input)
	//head := findHead(devices)
	queue := [][]string{{"you"}}
	allPaths := [][]string{}

	traverseSVR(&queue, &allPaths, devices)
	uniquePaths := make(map[string]struct{})

	for _, path := range allPaths {
		key := strings.Join(path, "->")
		uniquePaths[key] = struct{}{}
	}
	fmt.Println(len(uniquePaths))
}

func traverseSVR(queue *[][]string, allPaths *[][]string, devices map[string][]string) {
	for len(*queue) > 0 {
		curPath := (*queue)[0]
		*queue = (*queue)[1:]

		lastNode := curPath[len(curPath)-1]

		if lastNode == "out" {
			*allPaths = append(*allPaths, curPath)
			continue
		}

		if limbs, exists := devices[lastNode]; exists {
			for _, limb := range limbs {
				newPath := make([]string, len(curPath))
				copy(newPath, curPath)
				newPath = append(newPath, limb)
				*queue = append(*queue, newPath)
			}
		}
	}
}

func parseDevices(input []string) map[string][]string {
	res := make(map[string][]string)
	for i := 0; i < len(input); i += 2 {
		res[input[i]] = strings.Split(strings.Trim(input[i+1], " "), " ")
	}
	return res
}
