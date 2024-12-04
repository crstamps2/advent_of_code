package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
	"bufio"
	"math"
	"slices"
)

func main() {
	file, err := os.Open("inputs.txt")
	if err != nil {
			panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var leftList, rightList []float64
	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		left, _ := strconv.ParseFloat(line[0], 64)
		leftList = append(leftList, left)
		right, _ := strconv.ParseFloat(line[1], 64)
		rightList = append(rightList, right)
	}
	
	fmt.Printf("Your answer is: %v\n", FindCombinedDistance(leftList, rightList))
}

func FindCombinedDistance(list1 []float64, list2 []float64) int{
	slices.Sort(list1)
	slices.Sort(list2)
	var combinedDistance int
	for i := 0; i < len(list1); i++{
		combinedDistance += int(math.Abs(list1[i] - list2[i]))
	}
	return combinedDistance
}
