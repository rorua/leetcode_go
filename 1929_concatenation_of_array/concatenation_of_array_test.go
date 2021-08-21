package main

import "testing"

func TestGetConcatenation(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{name: "test1", input: []int{1, 2, 3}, expected: []int{1, 2, 3, 1, 2, 3}},
		{name: "test1", input: []int{1, 3, 2, 1}, expected: []int{1, 3, 2, 1, 1, 3, 2, 1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := getConcatenation(tt.input)
			for k, v := range output {
				if v != tt.expected[k] {
					t.Fatalf("getConcatenation  = %v, expected = %v", output, tt.expected)
				}
			}
		})
	}

}
