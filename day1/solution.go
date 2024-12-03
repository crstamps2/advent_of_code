package day1

import (
	"slices"
	"math"
)

func FindCombinedDistance(list1 []float64, list2 []float64) int{
	slices.Sort(list1)
	slices.Sort(list2)
	var combinedDistance int
	for i := 0; i < len(list1); i++{
		combinedDistance += int(math.Abs(list1[i] - list2[i]))
	}
	return combinedDistance
}
