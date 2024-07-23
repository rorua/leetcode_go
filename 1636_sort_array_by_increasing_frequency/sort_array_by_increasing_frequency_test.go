package sort_array_by_increasing_frequency

import "testing"

var tests = []struct {
	name     string
	input    []int
	expected []int
}{
	{
		name:     "test1",
		input:    []int{1, 1, 2, 2, 2, 3},
		expected: []int{3, 1, 1, 2, 2, 2},
	}, {
		name:     "test2",
		input:    []int{2, 3, 1, 3, 2},
		expected: []int{1, 3, 3, 2, 2},
	}, {
		name:     "test3",
		input:    []int{-1, 1, -6, 4, 5, -6, 1, 4, 1},
		expected: []int{5, -1, 4, 4, -6, -6, 1, 1, 1},
	},
}

func TestSortPeople(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := frequencySort(tt.input)
			for k, v := range output {
				if v != tt.expected[k] {
					t.Fatalf("Test %s failed: frequencySort(%v) = %v, expected = %v", tt.name, tt.input, output, tt.expected)
				}
			}
		})
	}
}
