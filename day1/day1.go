package day1

import (
	"slices"
)

func FindCombinedDistance(list1 []int, list2 []int) int{
	slices.Sort(list1)
	slices.Sort(list2)
	var combinedDistance int
	for i, v := range list1 {
		combinedDistance += v - list2[i]
	}
	return max(combinedDistance, -combinedDistance)
}
