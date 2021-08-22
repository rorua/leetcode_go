package running_sum_of_1d_array

import "testing"

func TestRunningSum(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{name: "test1", input: []int{1,2,3,4}, expected: []int{1,3,6,10}},
		{name: "test2", input: []int{1,1,1,1,1}, expected: []int{1,2,3,4,5}},
		{name: "test3", input: []int{3,1,2,10,1}, expected: []int{3,4,6,16,17}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := runningSum(tt.input)
			for k, v := range output {
				if tt.expected[k] != v {
					t.Errorf("runningSum = %v, want %v", output, tt.expected)
				}
			}
		})
	}

}
