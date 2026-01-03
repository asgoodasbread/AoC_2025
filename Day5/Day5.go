package main

import (
	"adventofcode2025/utils"
	"fmt"
	"sort"
	"strings"
)

type Interval struct {
	low  int
	high int
}

func main() {
	fmt.Println("Day 5 !")
	//part_1()
	part_2()
}

func part_2() {
	id_ranges, _ := utils.ReadDay5("Day5/day5_input.txt")
	//Parse id ranges as intervals
	id_range_intervals := make([]Interval, 0)
	for _, ran := range id_ranges {
		parts := strings.Split(ran, "-")
		myInterval := Interval{
			low:  utils.Atoi(parts[0]),
			high: utils.Atoi(parts[1]),
		}
		id_range_intervals = append(id_range_intervals, myInterval)
	}
	// Sort the intervals
	sort.Slice(id_range_intervals, func(i, j int) bool {
		return id_range_intervals[i].low < id_range_intervals[j].low
	})
	// Merge overlapping intervals
	merged_intervals := []Interval{}
	current := id_range_intervals[0]
	for i := 0; i < len(id_range_intervals); i++ {
		if id_range_intervals[i].low <= current.high {
			if id_range_intervals[i].high > current.high {
				current.high = id_range_intervals[i].high
			}
		} else {
			merged_intervals = append(merged_intervals, current)
			current = id_range_intervals[i]
		}
	}
	merged_intervals = append(merged_intervals, current)
	//fmt.Println(merged_intervals)
	res := 0
	for _, id_range := range merged_intervals {
		res += id_range.high - id_range.low + 1
	}
	fmt.Println(res)
}

func part_1() {
	id_ranges, ids := utils.ReadDay5("Day5/day5_input.txt")
	//Parse id ranges as intervals
	id_range_intervals := make([]Interval, 0, len(id_ranges))
	res := 0
	for _, ran := range id_ranges {
		parts := strings.Split(ran, "-")
		myInterval := Interval{
			low:  utils.Atoi(parts[0]),
			high: utils.Atoi(parts[1]),
		}
		id_range_intervals = append(id_range_intervals, myInterval)
	}
	for _, id := range ids {
		for _, id_range := range id_range_intervals {
			if id >= id_range.low && id <= id_range.high {
				res++
				break
			}
		}
	}
	fmt.Println(res)
}
