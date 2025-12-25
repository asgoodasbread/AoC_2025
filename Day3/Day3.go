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
	input := utils.ReadInputSep("Day3/test_input.txt", ",")
	digs := make([]int, 12)
	for _, x := range input {
		for i := 0; i < 12; i++ {
			findMax(x, i, digs, i)
		}
		fmt.Println(digs)
	}
}

func findMax(x string, maxIndex int, digs []int, curDig int) {
	if maxIndex == len(x) {
		return
	}
	max := -1
	for i := maxIndex; i < len(x)-1; i++ {
		digit := int(x[i] - '0')
		if digit > max {
			max = digit
		}
	}
	digs[curDig] = max

}
