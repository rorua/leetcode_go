package sort_the_people

import "testing"

var tests = []struct {
	name  string
	input struct {
		names   []string
		heights []int
	}
	expected []string
}{
	{
		name: "test1",
		input: struct {
			names   []string
			heights []int
		}{
			names:   []string{"Alex", "Ben", "Charlie", "David"},
			heights: []int{170, 173, 165, 190},
		},
		expected: []string{"David", "Ben", "Alex", "Charlie"},
	}, {
		name: "test2",
		input: struct {
			names   []string
			heights []int
		}{
			names:   []string{"Alice", "Bob", "Bob"},
			heights: []int{155, 185, 150},
		},
		expected: []string{"Bob", "Alice", "Bob"},
	},
}

func TestSortPeople(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := sortPeople2(tt.input.names, tt.input.heights)
			for k, v := range output {
				if v != tt.expected[k] {
					t.Fatalf("Test %s failed: sortPeople(%v, %v) = %v, expected = %v", tt.name, tt.input.names, tt.input.heights, output, tt.expected)
				}
			}
		})
	}
}

func BenchmarkSortPeople(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sortPeople2(tests[1].input.names, tests[1].input.heights)
	}
}

//BenchmarkSortPeople-12    	 5399250	       221.9 ns/op
//BenchmarkSortPeople-12    	 4964566	       219.4 ns/op
//BenchmarkSortPeople-12    	 5206838	       220.7 ns/op

//BenchmarkSortPeople-12    	12906358	        78.82 ns/op
