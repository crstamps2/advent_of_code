package main

import (
    "fmt"
    "regexp"
    "strconv"
)

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

func main() {
    input := "mul(123,456) some text mul(78,90) more text mul(1,2)"
    pairs := extractPairs(input)
    fmt.Println(pairs) // Output: [[123 456] [78 90] [1 2]]
}
