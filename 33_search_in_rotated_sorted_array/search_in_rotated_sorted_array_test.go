package search_in_rotated_sorted_array

import "testing"

func Test_search(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		target   int
		expected int
	}{
		{name: "test1", input: []int{4, 5, 6, 7, 0, 1, 2}, target: 0, expected: 4},
		{name: "test2", input: []int{4, 5, 6, 7, 0, 1, 2}, target: 3, expected: -1},
		{name: "test3", input: []int{1}, target: 0, expected: -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := search(tt.input, tt.target)
			if output != tt.expected {
				t.Errorf("output = %d, want %d", output, tt.expected)
			}
		})
	}
}
