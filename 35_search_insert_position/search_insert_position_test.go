package search_insert_position

import (
	"testing"
)

func Test_searchInsert(t *testing.T) {

	tests := []struct {
		name  string
		input struct {
			nums   []int
			target int
		}
		expected int
	}{
		{name: "test1", input: struct {
			nums   []int
			target int
		}{nums: []int{1, 3, 5, 6}, target: 5}, expected: 2},
		{name: "test2", input: struct {
			nums   []int
			target int
		}{nums: []int{1, 3, 5, 6}, target: 2}, expected: 1},
		{name: "test3", input: struct {
			nums   []int
			target int
		}{nums: []int{1, 3, 5, 6}, target: 7}, expected: 4},
		{name: "test4", input: struct {
			nums   []int
			target int
		}{nums: []int{1, 3, 5, 6}, target: 0}, expected: 0},
		{name: "test5", input: struct {
			nums   []int
			target int
		}{nums: []int{1}, target: 0}, expected: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := searchInsert(tt.input.nums, tt.input.target)
			if output != tt.expected {
				t.Fatalf("expected %v, got %v", tt.expected, output)
			}
		})
	}
}
