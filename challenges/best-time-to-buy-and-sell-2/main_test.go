package main

import (
    "fmt"
    "testing"
)

func Test_maxProfit(t *testing.T) {
	tests := []struct {
		nums           []int
        expected       int
	}{
		{
			nums: 		[]int{7, 1, 5, 3, 6, 4},
			expected: 	7,
		},
		{
			nums: 		[]int{1, 2, 3, 4, 5},
			expected: 	4,
		},
		{
			nums: 		[]int{7, 6, 4, 3, 1},
			expected: 	0,
		},
		{
			nums: 		[]int{7, 6},
			expected: 	0,
		},
		{
			nums: 		[]int{7, 6, 4},
			expected: 	0,
		},
		{
			nums: 		[]int{7, 4, 6},
			expected: 	2,
		},
		{
			nums: 		[]int{1, 5},
			expected: 	4,
		},
		{
			nums: 		[]int{1},
			expected: 	0,
		},
		{
			nums: 		[]int{2, 4, 1},
			expected: 	2,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("Testing with expected=%d", tt.expected), func(t *testing.T) {
            // Call the function
            k := maxProfit(tt.nums)

            // Check if k is correct
            if k != tt.expected {
                t.Errorf("Expected %d, got %d", tt.expected, k)
            }
        })
	}
}
