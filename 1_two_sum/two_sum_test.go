package two_sum

import "testing"

func TestTwoSum(t *testing.T) {
	type input struct {
		nums   []int
		target int
	}
	tests := []struct {
		name     string
		input    input
		expected []int
	}{
		{name: "test1", input: input{nums: []int{2, 7, 11, 15}, target: 9}, expected: []int{0, 1}},
		{name: "test2", input: input{nums: []int{3, 2, 4}, target: 6}, expected: []int{1, 2}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := twoSum(tt.input.nums, tt.input.target)
			if output[0] != tt.expected[0] && output[1] != tt.expected[1] {
				t.Errorf("twoSum = %v, want %v", output, tt.expected)
			}
		})
	}

}
