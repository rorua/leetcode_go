package binary_search

import "testing"

func TestSearch(t *testing.T) {

	tests := []struct {
		name     string
		input    []int
		target   int
		expected int
	}{
		{name: "test 1", input: []int{-1, 0, 3, 5, 9, 12}, target: 9, expected: 4},
		{name: "test 2", input: []int{-1, 0, 3, 5, 9, 12}, target: 2, expected: -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := search(tt.input, tt.target)
			if output != tt.expected {
				t.Fatalf("expected %d, got %d", tt.expected, output)
			}
		})
	}

}
