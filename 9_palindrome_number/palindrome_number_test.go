package palindrome_number

import (
	"testing"
)

func TestIsPalindrome(t *testing.T) {

	tests := []struct {
		name     string
		input    int
		expected bool
	}{
		{name: "test1", input: 1221, expected: true},
		{name: "test2", input: -121, expected: false},
		{name: "test3", input: 10, expected: false},
		{name: "test4", input: -101, expected: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := isPalindrome(tt.input)
			if output != tt.expected {
				t.Errorf("isPalindrome = %v, want %v", output, tt.expected)
			}
		})
	}

}
