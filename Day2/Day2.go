package main

import (
	"adventofcode2025/utils"
	"fmt"
	"strconv"
	"strings"
)

type ID struct {
	lbound int
	hbound int
}

func main() {
	part_1()
	part_2()
}

func part_2() {
	input := utils.ReadInputSep("Day2/input_day2.txt", ",")
	IDs := parse_ids(input)
	//fmt.Println(IDs)
	res := 0
	for _, id := range IDs {
		for i := id.lbound; i <= id.hbound; i++ {
			si := strconv.Itoa(i)
			if is_repeated(si) {
				res += i
			}
		}
	}
	fmt.Println(res)
}

func is_repeated(s string) bool {
	doubleS := s + s
	return strings.Contains(doubleS[1:len(doubleS)-1], s)
}

func part_1() {
	input := utils.ReadInputSep("Day2/input_day2.txt", ",")
	IDs := parse_ids(input)
	//fmt.Println(IDs)
	res := 0
	for _, id := range IDs {
		for i := id.lbound; i <= id.hbound; i++ {
			cs := strconv.Itoa(i)
			//fmt.Println(cs)
			if len(cs)%2 == 0 {
				if cs[:len(cs)/2] == cs[len(cs)/2:] {
					res += i
				}
			}
		}
	}
	fmt.Println(res)
}

func parse_ids(input []string) []ID {
	var result []ID
	for _, elem := range input {
		if len(elem) < 2 {
			continue
		}
		myID := ID{
			lbound: utils.Atoi(strings.Split(elem, "-")[0]),
			hbound: utils.Atoi(strings.Split(elem, "-")[1]),
		}
		result = append(result, myID)
	}
	return result
}
