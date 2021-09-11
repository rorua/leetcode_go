package letter_combinations_of_a_phone_number

import (
	"leetcode/utils"
	"testing"
)

func Test_letterCombinations(t *testing.T) {

	tests := []struct {
		name     string
		input    string
		expected []string
	}{
		{name: "test1", input: "23", expected: []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}},
		{name: "test2", input: "", expected: []string{}},
		{name: "test3", input: "2", expected: []string{"a", "b", "c"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := letterCombinations(tt.input)
			if len(output) != len(tt.expected) {
				t.Fatalf("expected len %v, got %v", len(tt.expected), len(output))
			}

			for _, v := range output {
				if !utils.InArray(v, tt.expected) {
					t.Fatalf("%s is not in expected array", v)
				}
			}
		})
	}
}
