package product_of_array_except_self

import (
	"reflect"
	"testing"
)

func Test_productExceptSelf(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{name: "test1", input: []int{1, 2, 3, 4}, expected: []int{24, 12, 8, 6}},
		{name: "test2", input: []int{-1, 1, 0, -3, 3}, expected: []int{0, 0, 9, 0, 0}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := productExceptSelf(tt.input)
			if !reflect.DeepEqual(output, tt.expected) {
				t.Fatalf("expected %v, got %v", tt.expected, output)
			}
		})
	}

	//for _, tt := range tests {
	//	t.Run(tt.name, func(t *testing.T) {
	//		output := productExceptSelf2(tt.input)
	//		if !reflect.DeepEqual(output, tt.expected) {
	//			t.Fatalf("expected %v, got %v", tt.expected, output)
	//		}
	//	})
	//}

}
