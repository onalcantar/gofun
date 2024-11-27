package main

import (
    "fmt"
    "testing"
)

func Test_productExceptSelf(t *testing.T) {
	tests := []struct {
		nums           []int
        expected       []int
	}{
		{
			nums: 		[]int{1, 2, 3, 4},
			expected: 	[]int{24, 12, 8, 6},
		},
		{
			nums: 		[]int{-1, 1, 0, -3, 3},
			expected: 	[]int{0, 0, 9, 0, 0},
		},
		{
			nums: 		[]int{4, 3, 2, 1, 2},
			expected: 	[]int{12, 16, 24, 48, 24},
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("Testing with expected=%d", tt.expected), func(t *testing.T) {
            // Call the function
            k := productExceptSelf(tt.nums)

            // Check if k is correct
			for i := 0; i < len(k); i++ {
				if k[i] != tt.expected[i] {
					t.Errorf("Expected %d, got %d", k[i], tt.expected[i])
				}
			}
        })
	}
}
