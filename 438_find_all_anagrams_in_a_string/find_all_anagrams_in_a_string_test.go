package find_all_anagrams_in_a_string

import "testing"

func TestFirstUniqChar(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		target   string
		expected []int
	}{
		{name: "test 1", input: "cbaebabacd", target: "abc", expected: []int{0, 6}},
		{name: "test 2", input: "abab", target: "ab", expected: []int{0, 1, 2}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := findAnagrams(tt.input, tt.target)
			if len(output)!= len(tt.expected) {
				t.Fatalf("expected len %v, got %v", len(tt.expected), len(output))
			}

			for i, v := range tt.expected {
				if v != output[i] {
					t.Fatalf("expected len %v, got %v", tt.expected, output)
				}
			}
		})
	}
}
