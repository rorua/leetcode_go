package maximum_product_subarray

import "testing"

func Test_maxProduct(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{name: "test1", input: []int{2,3,-2,4}, expected: 6},
		{name: "test2", input: []int{-2,0,-1}, expected: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := maxProduct(tt.input)
			if output != tt.expected {
				t.Errorf("reverse = %d, want %d", output, tt.expected)
			}
		})
	}

}
