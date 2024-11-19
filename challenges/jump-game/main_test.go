package main

import (
    "fmt"
    "testing"
)

func Test_canJump(t *testing.T) {
	tests := []struct {
		nums           []int
        expected       bool
	}{
		{
			nums: 		[]int{2, 5, 0, 0},
			expected: 	true,
		},
		{
			nums: 		[]int{2, 3, 1, 1, 4},
			expected: 	true,
		},
		{
			nums: 		[]int{3, 2, 1, 0, 4},
			expected: 	false,
		},
		{
			nums: 		[]int{2, 1, 1, 0},
			expected: 	true,
		},
		{
			nums: 		[]int{2, 1, 0, 0},
			expected: 	false,
		},
		{
			nums: 		[]int{5, 2},
			expected: 	true,
		},
		{
			nums: 		[]int{0, 2},
			expected: 	false,
		},
		{
			nums: 		[]int{0},
			expected: 	true,
		},
		{
			nums: 		[]int{},
			expected: 	true,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("Testing with expected=%t", tt.expected), func(t *testing.T) {
            // Call the function
            k := canJump(tt.nums)

            // Check if k is correct
            if k != tt.expected {
                t.Errorf("Expected %t, got %t", tt.expected, k)
            }
        })
	}
}
