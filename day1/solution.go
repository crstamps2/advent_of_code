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
	
	fmt.Printf("Combined Distance: %v\n", FindCombinedDistance(leftList, rightList))
	fmt.Printf("Similarity Score: %v\n", CalculateSimilarityScore(leftList, rightList))
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

func CalculateSimilarityScore(leftList []float64, rightList[]float64) int {
	var similarityScores []float64
	for i := 0; i < len(leftList); i++ {
		occurences := 0.0
		for x := 0; x < len(rightList); x++ {
			if leftList[i] == rightList[x] { occurences++ }
		}
		similarityScores = append(similarityScores, leftList[i]*occurences)
	}
	sum := 0.0
	for _, score := range similarityScores {
		sum += score
	}
	return int(sum)
}

