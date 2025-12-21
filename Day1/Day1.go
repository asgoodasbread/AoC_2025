package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type SafeCode struct {
	Direction int
	Value     int
}

func read_input() []string {
	var input []string
	file_name := "Day1/Day1_input.txt"
	file, err := os.Open(file_name)
	if err != nil {
		log.Fatalf("Error opening file!")
	}
	defer file.Close() //Defer make this line of code execute when the function returns
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		//fmt.Println(text)
		input = append(input, text)
	}
	return input
}

func main() {
	fmt.Println("Day1")
	//part_1()
	part_2()
}

func part_2() {
	input := read_input()
	var current_dial int = 50

	//Parsing safe codes
	var safeCodes []SafeCode
	for _, elem := range input {
		safeValue, err := strconv.Atoi(elem[1:])
		if err != nil {
			log.Fatalf("Error Converting string to integer")
			continue
		}

		//Encode direction
		dir := 0
		if elem[0] == 'L' {
			dir = 0
		} else {
			dir = 1
		}

		safeCode := SafeCode{
			Direction: dir,
			Value:     safeValue,
		}
		safeCodes = append(safeCodes, safeCode)
	}

	res := 0
	for _, code := range safeCodes {
		for i := 0; i < code.Value; i++ {
			if code.Direction == 0 {
				current_dial = (current_dial - 1 + 100) % 100
			} else {
				current_dial = (current_dial + 1) % 100
			}
			if current_dial == 0 {
				res += 1
			}
		}

	}
	fmt.Println("Current dial:", current_dial)
	fmt.Println("Hit-zero:", res)
}

func part_1() {
	input := read_input()
	var current_dial int = 50

	//Parsing safe codes
	var safeCodes []SafeCode
	for _, elem := range input {
		safeValue, err := strconv.Atoi(elem[1:])
		if err != nil {
			log.Fatalf("Error Converting string to integer")
			continue
		}

		//Encode direction
		dir := 0
		if elem[0] == 'L' {
			dir = 0
		} else {
			dir = 1
		}

		safeCode := SafeCode{
			Direction: dir,
			Value:     safeValue,
		}
		safeCodes = append(safeCodes, safeCode)
	}

	hit_zero := 0
	for _, code := range safeCodes {
		if code.Direction == 0 {
			current_dial -= code.Value
		} else {
			current_dial += code.Value
		}
		if current_dial < 0 {
			current_dial += 100
		}
		current_dial = current_dial % 100
		if current_dial == 0 {
			hit_zero += 1
		}
	}
	fmt.Println("Current dial:", current_dial)
	fmt.Println("Hit-zero:", hit_zero)
}
