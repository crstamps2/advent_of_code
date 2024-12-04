package main

import "testing"

func TestIsReportSafe(t *testing.T) {

	tests := []struct {
		name string
		readings []int
		want bool
	}{
		{"should be true", []int{1,2,3,4,5}, true},
		{"should be true", []int{5,4,3,2,1}, true},
		{"should be true", []int{1,3,5,7,9}, true},
		{"should be true", []int{9,7,5,3,1}, true},
		{"should be false", []int{1,2,6,7,8}, false},
		{"should be false", []int{8,7,6,2,1}, false},
		{"should be false", []int{1,2,3,3,4}, false},
		{"should be false", []int{4,3,3,2,1}, false},
		{"should be false", []int{1,2,3,2,1}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsReportSafe(tt.readings)
			if result != tt.want {
				t.Errorf("inputs: %v , got: %v, wanted: %v", tt.readings, result, tt.want)
			} 
		})
	}
}
