package find_the_city_with_the_smallest_number_of_neighbors_at_a_threshold_distance

import "testing"

func TestFindTheCity(t *testing.T) {
	tests := []struct {
		name  string
		input struct {
			n                 int
			edges             [][]int
			distanceThreshold int
		}
		expected int
	}{
		{
			name: "test1",
			input: struct {
				n                 int
				edges             [][]int
				distanceThreshold int
			}{
				n: 4,
				edges: [][]int{
					{0, 1, 3},
					{1, 2, 1},
					{1, 3, 4},
					{2, 3, 1},
				},
				distanceThreshold: 4,
			},
			expected: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := findTheCity(tt.input.n, tt.input.edges, tt.input.distanceThreshold)
			if output != tt.expected {
				t.Fatalf("findTheCity  = %v, expected = %v", output, tt.expected)
			}
		})
	}
}
