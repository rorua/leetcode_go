package divide_array_into_equal_pairs

import "testing"

var tests = []struct {
	name     string
	input    []int
	expected bool
}{
	{
		name:     "1",
		input:    []int{1, 1, 2, 2, 3, 3},
		expected: true,
	},
	{
		name:     "2",
		input:    []int{1, 2, 3, 4},
		expected: false,
	},
}

func TestDivideArray(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := divideArray(tt.input)
			if output != tt.expected {
				t.Fatalf("Test %s failed: divideArray(%v) = %v, expected = %v", tt.name, tt.input, output, tt.expected)
			}
		})
	}
}

func BenchmarkDivideArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		divideArray(tests[0].input)
	}
}

//BenchmarkDivideArray-12    	 8534628	       141.3 ns/op
//BenchmarkDivideArray-12    	 6667857	       180.0 ns/op
//BenchmarkDivideArray-12    	 6615084	       179.9 ns/op
