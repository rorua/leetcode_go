package find_min_in_rotated_sorted_array

import "testing"

func Test_findMin(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{name: "test1", input: []int{2, 3, -2, -1}, expected: -2},
		{name: "test2", input: []int{3, 4, 5, 1, 2}, expected: 1},
		{name: "test3", input: []int{4, 5, 6, 7, 0, 1, 2}, expected: 0},
		{name: "test4", input: []int{11, 13, 15, 17}, expected: 11},
		{name: "test5", input: []int{2, 1}, expected: 1},
		{name: "test6", input: []int{2}, expected: 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := findMin(tt.input)
			if output != tt.expected {
				t.Errorf("reverse = %d, want %d", output, tt.expected)
			}
		})
	}

}
