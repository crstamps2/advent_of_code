package day1

import (
	"slices"
)

func FindCombinedDistance(list1 []float64, list2 []float64) float64{
	slices.Sort(list1)
	slices.Sort(list2)
	var combinedDistance float64
	for i, v := range list1 {
		combinedDistance += v - list2[i]
	}
	return max(combinedDistance, -combinedDistance)
}
