package main

import (
	"fmt"
	"testing"
)

func Test_plusOne(t *testing.T) {
	tests := []struct {
		digits         []int
        expected       []int
	}{
			{
				digits: 	[]int{1, 2, 3},
				expected: 	[]int{1, 2, 4},
			},
			{
				digits: 	[]int{4, 3, 2, 1},
				expected: 	[]int{4, 3, 2, 2},
			},
			{
				digits: 	[]int{9},
				expected: 	[]int{1, 0},
			},
			{
				digits: 	[]int{9, 9, 9},
				expected: 	[]int{1, 0, 0, 0},
			},
			{
				digits: 	[]int{0},
				expected: 	[]int{1},
			},
			{
				digits: 	[]int{1, 9, 9},
				expected: 	[]int{2, 0, 0},
			},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("Testing with expected=%v", tt.expected), func(t *testing.T) {
            // Call the function
            k := plusOne(tt.digits)

            // Check if k is correct
			for i, v := range k {
				if v != tt.expected[i] {
					t.Errorf("Expected %d, got %d", tt.expected, v)
				}
			}
        })
	}
}
