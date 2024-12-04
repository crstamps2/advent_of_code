package main

import (
	// "fmt"
	// "os"
	// "strings"
	// "strconv"
	// "bufio"
	"math"
	// "slices"
)

func main() {

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
