package main

import (
    "fmt"
    "testing"
)

func Test_majorityElement(t *testing.T) {
	tests := []struct {
		nums           []int
        expected       int
	}{
		{
			nums: []int{3, 2, 3},
			expected: 3,
		},
		{
			nums: []int{2, 2, 1, 1, 1, 2, 2},
			expected: 2,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("Testing with expected=%d", tt.expected), func(t *testing.T) {
            // Call the function
            k := majorityElement(tt.nums)

            // Check if k is correct
            if k != tt.expected {
                t.Errorf("Expected %d, got %d", tt.expected, k)
            }
        })
	}
}
