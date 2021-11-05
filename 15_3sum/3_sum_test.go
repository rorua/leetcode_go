package three_sum

import (
	"reflect"
	"testing"
)

func Test_threeSum(t *testing.T) {

	tests := []struct {
		name     string
		input    []int
		expected [][]int
	}{
		{name: "test1", input: []int{-1, 0, 1, 2, -1, -4}, expected: [][]int{{-1, -1, 2}, {-1, 0, 1}}},
		{name: "test2", input: []int{0}, expected: [][]int{}},
		{name: "test3", input: []int{}, expected: [][]int{}},
		{name: "test4", input: []int{-4, -2, -2, -2, 0, 1, 2, 2, 2, 3, 3, 4, 4, 6, 6}, expected: [][]int{{-4, -2, 6}, {-4, 0, 4}, {-4, 1, 3}, {-4, 2, 2}, {-2, -2, 4}, {-2, 0, 2}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := threeSum(tt.input)
			if len(output) != len(tt.expected) {
				t.Fatalf("expected len %v, got %v", len(tt.expected), len(output))
			}
			if !reflect.DeepEqual(output, tt.expected) {
				t.Fatalf("expected %v, got %v", tt.expected, output)
			}
		})
	}
}
