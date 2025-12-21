package main

import (
	"adventofcode2025/utils"
	"fmt"
	"strings"
)

type ID struct {
	lbound int
	hbound int
}

func main() {
	part_1()
	//part_2()
}

func part_1() {
	input := utils.ReadInputSep("Day2/test_day2.txt", ",")
	IDs := parse_ids(input)
	fmt.Println(IDs)

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
