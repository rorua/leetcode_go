package sort_an_array

import "testing"

var tests = []struct {
	name     string
	input    []int
	expected []int
}{
	{
		name:     "1",
		input:    []int{5, 2, 3, 1},
		expected: []int{1, 2, 3, 5},
	},
	{
		name:     "2",
		input:    []int{5, 1, 1, 2, 0, 0},
		expected: []int{0, 0, 1, 1, 2, 5},
	},
}

func TestSortArray(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := tt.input
			output := sortArray(tt.input)
			for k, v := range output {
				if v != tt.expected[k] {
					t.Fatalf("Test %s failed: sortArray(%v) = %v, expected = %v", tt.name, input, output, tt.expected)
				}
			}
		})
	}
}
