package reverse_integer

import "testing"

func TestReverse(t *testing.T) {

	tests := []struct {
		name     string
		input    int
		expected int
	}{
		{name: "test1", input: 123, expected: 321},
		{name: "test2", input: 4340, expected: 434},
		{name: "test3", input: 123456789, expected: 987654321},
		{name: "test4", input: -2147483648, expected: 0},
		{name: "test5", input: 1534236469, expected: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := reverse(tt.input)
			if output != tt.expected {
				t.Errorf("reverse = %d, want %d", output, tt.expected)
			}
		})
	}

}
