package day1

import "testing"

func TestFindCombinedDistance(t *testing.T) {
	list1 := []int{5,1,7,6}
	list2 := []int{9,7,5,2}

	result := FindCombinedDistance(list1, list2)

	if result != 4 {
		t.Errorf("got: %v, wanted: %v", result, 4)
	} 
}
