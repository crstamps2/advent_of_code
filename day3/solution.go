package main

import (
	"os"
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	file, err := os.ReadFile("inputs.txt")
	if err != nil {
			panic(err)
	}

	str := string(file)
	pairs := extractPairs(str)

	var total int
	for _, pair := range pairs {
		total += mul(pair)
	}

	fmt.Printf("total of mul pairs: %v\n", total)
}

func extractPairs(input string) [][2]int {
	// Compile the regex pattern
	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

	// Find all matches
	matches := re.FindAllStringSubmatch(input, -1)

	// Prepare the result slice
	var result [][2]int

	// Iterate over the matches and extract the numbers
	for _, match := range matches {
		if len(match) == 3 {
			// Convert the extracted strings to integers
			num1, err1 := strconv.Atoi(match[1])
			num2, err2 := strconv.Atoi(match[2])
			if err1 == nil && err2 == nil {
				// Append the pair to the result slice
				result = append(result, [2]int{num1, num2})
			}
		}
	}

	return result
}

func mul(pair [2]int) int {
	return pair[0] * pair[1]
}
