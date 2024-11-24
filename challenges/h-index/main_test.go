package main

import (
    "fmt"
    "testing"
)

func Test_hIndex(t *testing.T) {
	tests := []struct {
		nums           []int
        expected       int
	}{
		{
			nums: 		[]int{3, 0, 6, 1, 5},
			expected: 	3,
		},
		{
			nums: 		[]int{1, 3, 1},
			expected: 	1,
		},
		{
			nums: 		[]int{100},
			expected: 	1,
		},
		{
			nums: 		[]int{},
			expected: 	0,
		},
		{
			nums: 		[]int{10, 8, 5, 4, 3},
			expected: 	4,
		},
		{
			nums: 		[]int{100, 50, 25, 10, 1},
			expected: 	4,
		},
		{
			nums: 		[]int{0, 0, 0, 0, 0},
			expected: 	0,
		},
		{
			nums: 		[]int{-1, -2, 3, 4, 5, 6},
			expected: 	3,
		},
		{
			nums: 		[]int{-1, -2, -3},
			expected: 	0,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("Testing with expected=%d", tt.expected), func(t *testing.T) {
            // Call the function
            k := hIndex(tt.nums)

            // Check if k is correct
            if k != tt.expected {
                t.Errorf("Expected %d, got %d", tt.expected, k)
            }
        })
	}
}
