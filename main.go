package main

import (
	"fmt"
	"encoding/csv"
	"os"
	"strconv"
	"codystamps.dev/day1"
)

func main() {
		f1, err := os.Open("day1/list1.csv")
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

		f2, err := os.Open("day1/list2.csv")
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
		
    fmt.Printf("Your answer is: %v\n", day1.FindCombinedDistance(list1, list2))
}
