package main

import (
    "fmt"
    "testing"
)

func Test_canJump(t *testing.T) {
	tests := []struct {
		nums           []int
        expected       int
	}{
		{
			nums: 		[]int{2, 3, 1, 1, 4},
			expected: 	2,
		},
		{
			nums: 		[]int{2, 3, 0, 1, 4},
			expected: 	2,
		},
		{
			nums: 		[]int{2, 1, 3, 1, 4},
			expected: 	2,
		},
		{
			nums: 		[]int{0},
			expected: 	0,
		},
		{
			nums: 		[]int{1, 2},
			expected: 	1,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("Testing with expected=%d", tt.expected), func(t *testing.T) {
            // Call the function
            k := jump(tt.nums)

            // Check if k is correct
            if k != tt.expected {
                t.Errorf("Expected %d, got %d", tt.expected, k)
            }
        })
	}
}
