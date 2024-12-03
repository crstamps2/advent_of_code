package main

import (
	"fmt"
	"encoding/csv"
	"os"
	"strconv"
	"slices"
	"math"
)

func main() {
	f1, err := os.Open("list1.csv")
	if err != nil {
			panic(err)
	}
	defer f1.Close()

	csvReader := csv.NewReader(f1)
	records, err := csvReader.ReadAll()
	var list1 []float64
	for index, record := range records {
		i, _ := strconv.ParseFloat(record[index], 64)
		list1 = append(list1, i)
	}
	if err != nil {
			panic(err)
	}

	f2, err := os.Open("list2.csv")
	if err != nil {
			panic(err)
	}
	defer f2.Close()

	csvReader = csv.NewReader(f2)
	records, err = csvReader.ReadAll()
	var list2 []float64
	for index, record := range records {
		i, _ := strconv.ParseFloat(record[index], 64)
		list2 = append(list2, i)
	}
	if err != nil {
			panic(err)
	}
	
	fmt.Printf("Your answer is: %v\n", FindCombinedDistance(list1, list2))
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
