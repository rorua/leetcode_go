package roman_to_integer

import "testing"

func TestRomanToInt(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{name: "test1", input: "X", expected: 10},
		{name: "test2", input: "IV", expected: 4},
		{name: "test3", input: "XXI", expected: 21},
		{name: "test4", input: "MCMXCIV", expected: 1994},
		{name: "test5", input: "LVIII", expected: 58},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := romanToInt(tt.input)
			if output != tt.expected {
				t.Errorf("romanToInt = %d, want %d", output, tt.expected)
			}
		})
	}
}
