package main

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
		{"should be 20", []float64{9,8,7,6}, []float64{11,12,13,14}, 20},
		{"should also be 20", []float64{11,12,13,14}, []float64{9,8,7,6}, 20},
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

func TestCalculateSimilarityScore(t *testing.T) {
	tests := []struct {
		name string
		list1 []float64
		list2 []float64
		want int
	}{
		{"should be 12", []float64{5,1,7,6}, []float64{9,7,5,2}, 12},
		{"should also be 12", []float64{9,7,5,2}, []float64{5,1,7,6}, 12},
		{"should be 26", []float64{9,7,7,6}, []float64{11,6,6,7}, 26},
		{"should be 23", []float64{1,2,2,3,3}, []float64{1,3,3,2,3}, 23},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CalculateSimilarityScore(tt.list1, tt.list2)
			if result != tt.want {
				t.Errorf("got: %v, wanted: %v", result, tt.want)
			} 
		})
	}
}
