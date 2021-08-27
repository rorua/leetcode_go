package first_unique_character_in_a_string

import "testing"

func TestName(t *testing.T) {

	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{name: "test 1", input: "leetcode", expected: 0},
		{name: "test 2", input: "loveleetcode", expected: 2},
		{name: "test 3", input: "aabb", expected: -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := firstUniqChar(tt.input)
			if output != tt.expected {
				t.Fatalf("expected %v, got %v", tt.expected, output)
			}

		})
	}
}
