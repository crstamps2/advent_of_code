package day1

import "testing"

func TestFindCombinedDistance(t *testing.T) {
	tests := []struct {
		name string
		list1 []float64
		list2 []float64
		want int
	}{
		{"should be 4", []float64{5,1,7,6}, []float64{9,7,5,2}, 4},
		{"should also be 4", []float64{9,7,5,2}, []float64{5,1,7,6}, 4},
		{"should be 28", []float64{-5,1,7,6}, []float64{9,8,5,15}, 28},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FindCombinedDistance(tt.list1, tt.list2)
			if result != tt.want {
				t.Errorf("got: %v, wanted: %v", result, tt.want)
			} 
		})
	}
}
