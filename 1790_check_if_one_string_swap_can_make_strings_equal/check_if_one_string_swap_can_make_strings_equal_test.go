package check_if_one_string_swap_can_make_strings_equal

import "testing"

func TestAreAlmostEqual(t *testing.T) {

	tests := []struct {
		name     string
		input    []string
		expected bool
	}{
		{name: "test1", input: []string{"bank", "kanb"}, expected: true},
		{name: "test2", input: []string{"attack", "defend"}, expected: false},
		{name: "test3", input: []string{"kelb", "kelb"}, expected: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := areAlmostEqual(tt.input[0], tt.input[1])
			if output != tt.expected {
				t.Errorf("buddyStrings = %v, want %v", output, tt.expected)
			}
		})
	}
}