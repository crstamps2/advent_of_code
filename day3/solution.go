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
	pairs := extractDoAndDontPairs(str)

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

func extractDoAndDontPairs(input string) [][2]int {
	// Compile the regex patterns
	mulRe := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	doRe := regexp.MustCompile(`do\(\)`)
	dontRe := regexp.MustCompile(`don't\(\)`)

	// Prepare the result slice
	var result [][2]int

	// Track the current state (whether to include or exclude)
	include := true

	// Iterate through the string
	for len(input) > 0 {
			// Find the next occurrence of do(), don't(), or mul()
			doIndex := doRe.FindStringIndex(input)
			dontIndex := dontRe.FindStringIndex(input)
			mulIndex := mulRe.FindStringIndex(input)

			// Determine the earliest occurrence
			nextIndex := len(input)
			if doIndex != nil && doIndex[0] < nextIndex {
					nextIndex = doIndex[0]
			}
			if dontIndex != nil && dontIndex[0] < nextIndex {
					nextIndex = dontIndex[0]
			}
			if mulIndex != nil && mulIndex[0] < nextIndex {
					nextIndex = mulIndex[0]
			}

			// Process the earliest occurrence
			if nextIndex == len(input) {
					break
			} else if doIndex != nil && doIndex[0] == nextIndex {
					include = true
					input = input[doIndex[1]:]
			} else if dontIndex != nil && dontIndex[0] == nextIndex {
					include = false
					input = input[dontIndex[1]:]
			} else if mulIndex != nil && mulIndex[0] == nextIndex {
					if include {
							match := mulRe.FindStringSubmatch(input)
							if len(match) == 3 {
									num1, err1 := strconv.Atoi(match[1])
									num2, err2 := strconv.Atoi(match[2])
									if err1 == nil && err2 == nil {
											result = append(result, [2]int{num1, num2})
									}
							}
					}
					input = input[mulIndex[1]:]
			} else {
					break
			}
	}

	return result
}

func mul(pair [2]int) int {
	return pair[0] * pair[1]
}
