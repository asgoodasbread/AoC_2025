package utils

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func ReadInputSep(file_name string, separator string) []string {
	content, err := os.ReadFile(file_name)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	str := string(content)

	str = strings.ReplaceAll(str, "\n", separator)
	str = strings.ReplaceAll(str, "\r", "")

	parts := strings.Split(str, separator)
	return parts
}

func Atoi(s string) int {
	val, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("Error converting '%s' to int: %v", s, err)
	}
	return val
}

// func main() {
// 	input := ReadInputSep("Day2/test_day2.txt", ",")
// 	//fmt.Println([]rune(input[5]))
// 	for i, x := range input {
// 		if i == 7 {
// 			fmt.Println()
// 		}
// 		//fmt.Println(x)
// 	}
// }
