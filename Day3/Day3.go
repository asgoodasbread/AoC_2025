package main

import (
	"adventofcode2025/utils"
	"fmt"
)

func main() {
	fmt.Println("Day 3")
	//part_1()
	part_2()
}

func part_1() {
	input := utils.ReadInputSep("Day3/day3_input.txt", ",")

	res := 0
	for _, x := range input {
		max1, max2 := getSec(x)
		res += max1*10 + max2
		//fmt.Println(max1, max2)
	}
	fmt.Println(res)
}

func getMax(x string) (int, int) {
	max := -1
	maxIndex := -1
	for i := 0; i < len(x)-1; i++ {
		digit := int(x[i] - '0')
		if digit > max {
			max = digit
			maxIndex = i
		}
	}
	return max, maxIndex
}

func getSec(x string) (int, int) {
	max1, maxIndex := getMax(x)
	max2 := -1
	for i := maxIndex + 1; i < len(x); i++ {
		digit := int(x[i] - '0')
		if digit > max2 {
			max2 = digit
		}
	}
	return max1, max2
}

func part_2() {
	input := utils.ReadInputSep("Day3/day3_input.txt", ",")
	res := 0
	for _, x := range input {
		nums := make([]int, 0, len(x))
		nums = get_nums(nums, x)
		//fmt.Println(nums)
		stack := make([]int, 0, 12)
		skips := len(nums) - 12
		for _, current := range nums {
			for len(stack) > 0 && current > stack[len(stack)-1] && skips > 0 {
				//pop the top
				stack = stack[:len(stack)-1]
				skips--
			}
			stack = append(stack, current)
		}
		//fmt.Println(stack[:len(stack)-skips])
		res += convertToInt(stack[:len(stack)-skips])
	}
	fmt.Println(res)
}

func get_nums(stack []int, s string) []int {
	for _, x := range s {
		stack = append(stack, int(x-'0'))
	}
	return stack
}

func convertToInt(x []int) int {
	res := 0
	for _, elem := range x {
		res = (res * 10) + elem
	}
	return res
}
