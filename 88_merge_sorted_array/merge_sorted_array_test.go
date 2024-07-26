package merge_sorted_array

import (
	"testing"
)

var tests = []struct {
	name  string
	input struct {
		num1 []int
		num2 []int
		m    int
		n    int
	}
	expected []int
}{
	{
		name: "1",
		input: struct {
			num1 []int
			num2 []int
			m    int
			n    int
		}{
			num1: []int{1, 2, 3, 0, 0, 0},
			num2: []int{2, 5, 6},
			m:    3,
			n:    3,
		},
		expected: []int{1, 2, 2, 3, 5, 6},
	}, {
		name: "2",
		input: struct {
			num1 []int
			num2 []int
			m    int
			n    int
		}{
			num1: []int{1},
			num2: []int{},
			m:    1,
			n:    0,
		},
		expected: []int{1},
	}, {
		name: "3",
		input: struct {
			num1 []int
			num2 []int
			m    int
			n    int
		}{
			num1: []int{0},
			num2: []int{1},
			m:    0,
			n:    1,
		},
		expected: []int{1},
	},
}

func TestMerge(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			merge(tt.input.num1, tt.input.m, tt.input.num2, tt.input.n)
			for k, v := range tt.input.num1 {
				if v != tt.expected[k] {
					t.Fatalf("Test %s failed: merge() = %v, expected = %v", tt.name, tt.input.num1, tt.expected)
				}
			}
		})
	}
}
