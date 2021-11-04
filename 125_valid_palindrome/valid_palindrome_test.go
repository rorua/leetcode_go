package valid_palindrome

import "testing"

func Test_isPalindrome(t *testing.T) {

	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{name: "test1", input: "A man, a plan, a canal: Panama", expected: true},
		{name: "test2", input: "race a car", expected: false},
		{name: "test3", input: " ", expected: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := isPalindrome(tt.input)
			if output != tt.expected {
				t.Fatalf("expected %v, got %v", tt.expected, output)
			}
		})
	}
}
