package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
	"bufio"
	"math"
)

func main() {
	file, err := os.Open("inputs.txt")
	if err != nil {
			panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var stringReadings []string
	var reports [][]int
	for scanner.Scan() {
		stringReadings = strings.Fields(scanner.Text())
		var readings []int
		for _, v := range stringReadings {
			i, _ := strconv.Atoi(v)
			readings = append(readings, i)
		}
		reports = append(reports, readings)
	}
	totalSafe := 0
	for _, report := range reports {
		if IsReportSafe(report) { totalSafe++ }
	}
	fmt.Printf("Number of Safe Readings: %v\n", totalSafe)
}

func IsReportSafe(readings []int) bool {
	var readingWentUp bool
	for index, reading := range readings {
		if index == 0 { continue } //skip first number
		if reading == (readings[index-1]) { return false }
		if math.Abs(float64(reading - readings[index-1])) > 3 { return false }

		if reading > readings[index-1] { readingWentUp = true }
		if reading > readings[index-1] && readingWentUp {
			continue
		} else if reading < readings[index-1] && !readingWentUp {
			continue
		} else {
			return false
		}
	}
	return true
}
