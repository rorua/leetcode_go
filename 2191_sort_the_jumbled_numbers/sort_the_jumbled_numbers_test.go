package sort_the_jumbled_numbers

import "testing"

var tests = []struct {
	name  string
	input struct {
		mapping []int
		nums    []int
	}
	expected []int
}{
	{
		name: "1",
		input: struct {
			mapping []int
			nums    []int
		}{mapping: []int{8, 9, 4, 0, 2, 1, 3, 5, 7, 6}, nums: []int{991, 338, 38}},
		expected: []int{338, 38, 991},
	}, {
		name: "2",
		input: struct {
			mapping []int
			nums    []int
		}{mapping: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, nums: []int{789, 456, 123}},
		expected: []int{123, 456, 789},
	}, {
		name: "3",
		input: struct {
			mapping []int
			nums    []int
		}{
			mapping: []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
			nums:    []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		expected: []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
	},
}

func TestSortJumbled(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := sortJumbled(tt.input.mapping, tt.input.nums)
			if len(output) != len(tt.expected) {
				t.Fatalf("Test %s failed: sortJumbled(%v) = %v, expected = %v len no queal", tt.name, tt.input, output, tt.expected)
			}

			for k, v := range output {
				if v != tt.expected[k] {
					t.Fatalf("Test %s failed: sortJumbled(%v) = %v, expected = %v", tt.name, tt.input, output, tt.expected)
				}
			}
		})
	}
}
