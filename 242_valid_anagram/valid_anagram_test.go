package valid_anagram

import "testing"

func TestIsAnagram(t *testing.T) {

	tests := []struct {
		name     string
		input    string
		target   string
		expected bool
	}{
		{name: "test 1", input: "foobar", target: "boofar", expected: true},
		{name: "test 1", input: "anagram", target: "nagaram", expected: true},
		{name: "test 1", input: "rat", target: "car", expected: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := isAnagram(tt.input, tt.target)
			if output != tt.expected {
				t.Fatalf("expected len %v, got %v", tt.expected, output)
			}

		})
	}
}
