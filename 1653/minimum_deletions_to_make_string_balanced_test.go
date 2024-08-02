package minimum_deletions_to_make_string_balanced

import "testing"

func TestFindTheCity(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "test1",
			input:    "aababbab",
			expected: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := minimumDeletions(tt.input)
			if output != tt.expected {
				t.Fatalf("minimumDeletions(%v)  = %d, expected = %d", tt.input, output, tt.expected)
			}
		})
	}
}
