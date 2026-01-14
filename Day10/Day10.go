package main

import (
	"adventofcode2025/utils"
	"fmt"
	"strings"

	"gonum.org/v1/gonum/stat/combin"
)

type Machine struct {
	lights  map[int]bool
	buttons [][]int
	joltage []int
}

func main() {
	//part_1()
	part_2()
}

func part_2() {
	input := utils.ReadDay10("Day10/test_input.txt")
	machines := parseInput(input)
	res := 0

	for _, machine := range machines {
		memo := make(map[string]int)
		res += minPresses(machine.buttons, machine.joltage, memo)
	}
	fmt.Println(res)
}

func key(target []int) string {
	return fmt.Sprint(target)
}

func minPresses(buttons [][]int, target []int, memo map[string]int) int {
	allZero := true
	for _, v := range target {
		if v != 0 {
			allZero = false
			break
		}
	}
	if allZero {
		return 0
	}
	k := key(target)
	if val, ok := memo[k]; ok {
		return val
	}
	n, m := len(buttons), len(target)
	limit := 1 << n
	best := -1
	for mask := range limit {
		remainder := make([]int, m)
		copy(remainder, target)
		costPhase1, poss := 0, true
		for b := range n {
			if (mask & (1 << b)) != 0 {
				costPhase1++
				for _, rowIdx := range buttons[b] {
					if rowIdx < len(remainder) {
						remainder[rowIdx]--
					}
				}
			}
		}
		for i := range m {
			if remainder[i] < 0 || remainder[i]%2 != 0 {
				poss = false
				break
			}
		}
		if poss {
			nextTarget := make([]int, m)
			for i := range m {
				nextTarget[i] = remainder[i] / 2
			}
			res := minPresses(buttons, nextTarget, memo)
			if res != -1 {
				totalCost := costPhase1 + 2*res
				if best == -1 || totalCost < best {
					best = totalCost
				}
			}
		}
	}
	memo[k] = best
	return best
}

func part_1() {
	input := utils.ReadDay10("Day10/day10_input.txt")

	machines := parseInput(input)
	res := 0

	for _, machine := range machines {
		found := false
		for i := 1; i < len(machine.buttons)+1; i++ {
			//Generate all possible i-sizes button combinations
			combsList := combin.Combinations(len(machine.buttons), i)
			for _, indices := range combsList {
				currentLights := make(map[int]bool)
				var currentCombination [][]int
				for _, idx := range indices {
					currentCombination = append(currentCombination, machine.buttons[idx])
				}
				//fmt.Println(currentCombination)
				for _, combination := range currentCombination {
					toggleLights(currentLights, combination)
				}
				if compareSets(machine.lights, currentLights) {
					res += i
					found = true
					break
				}
			}
			if found {
				break
			}
		}
	}
	fmt.Println(res)
}

func compareSets(a, b map[int]bool) bool {
	if len(a) != len(b) {
		return false
	}
	for k := range a {
		if !b[k] {
			return false
		}
	}
	return true
}

func toggleLights(lights map[int]bool, buttonIndices []int) {
	for _, idx := range buttonIndices {
		if lights[idx] {
			delete(lights, idx)
		} else {
			lights[idx] = true
		}
	}
}

func parseInput(input [][]string) []Machine {
	machines := make([]Machine, 0, len(input))

	for _, x := range input {
		myMachine := Machine{
			lights:  parseLights(x[0]),
			buttons: parseButtons(x[1]),
			joltage: parseJoltage(x[2]),
		}
		machines = append(machines, myMachine)
	}
	return machines
}

func parseLights(line string) map[int]bool {
	res := make(map[int]bool)

	for i, char := range line {
		if char == '#' {
			res[i] = true
		}
	}

	return res
}

func parseJoltage(line string) []int {
	var res []int
	input := strings.Trim(line, "{}")
	ns := strings.Split(input, ",")
	for _, x := range ns {
		res = append(res, utils.Atoi(x))
	}
	return res
}

func parseButtons(line string) [][]int {
	var res [][]int
	input := strings.Split(line, ")")
	for _, x := range input {
		button := make([]int, 0)
		x = strings.Trim(x, " (")
		if x == "" {
			continue
		}
		y := strings.Split(x, ",")
		for _, a := range y {
			button = append(button, utils.Atoi(strings.TrimSpace(a)))
		}
		res = append(res, button)
	}
	//fmt.Println(input)
	return res
}
