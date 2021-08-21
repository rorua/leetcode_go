package longest_substring

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {

	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{name: "test1", input: "", expected: 0},
		{name: "test2", input: " ", expected: 1},
		{name: "test3", input: "tmmzuxt", expected: 5},
		{name: "test4", input: "abcabcbb", expected: 3},
		{name: "test5", input: "bbbbb", expected: 1},
		{name: "test6", input: "pwwkew", expected: 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := lengthOfLongestSubstring(tt.input)
			if output != tt.expected {
				t.Errorf("lengthOfLongestSubstring = %d, want %d", output, tt.expected)
			}
		})
	}
}
