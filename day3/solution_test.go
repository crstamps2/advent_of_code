package main


import (
    "reflect"
    "testing"
)

func TestExtractPairs(t *testing.T) {
    tests := []struct {
        input string
        want  [][2]int
    }{
        {"mul(123,456) some text mul(78,90) more text mul(1,2)", [][2]int{{123, 456}, {78, 90}, {1, 2}}},
        {"no matches here", nil},
        {"mul(12,34) and mul(56,78)", [][2]int{{12, 34}, {56, 78}}},
        {"mul(1,2) mul(3,4) mul(5,6)", [][2]int{{1, 2}, {3, 4}, {5, 6}}},
        {"mul(999,888) mul(777,666)", [][2]int{{999, 888}, {777, 666}}},
    }

    for _, tt := range tests {
        t.Run(tt.input, func(t *testing.T) {
            got := extractPairs(tt.input)
            if !reflect.DeepEqual(got, tt.want) {
                t.Errorf("extractPairs(%q) = %v, want %v", tt.input, got, tt.want)
            }
        })
    }
}
