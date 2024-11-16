package main

import (
    "fmt"
    "testing"
)

func Test_rotate(t *testing.T) {
	tests := []struct {
		nums			[]int
        numsToRotate	int
		expected 		[]int
	}{
		{
			nums: 			[]int{1, 2, 3, 4, 5, 6, 7},
			numsToRotate: 	3,
			expected: 		[]int{5, 6, 7, 1, 2, 3, 4},
		},
		{
			nums: 			[]int{-1, -100, 3, 99},
			numsToRotate: 	2,
			expected: 		[]int{3, 99, -1, -100},
		},
		{
			nums: 			[]int{-1, -100},
			numsToRotate: 	1,
			expected: 		[]int{-100, -1},
		},
		{
			nums: 			[]int{-1, -100},
			numsToRotate: 	0,
			expected: 		[]int{-1, -100},
		},
		{
			nums: 			[]int{-1, -100, 3, 5, 10, 7, 8},
			numsToRotate: 	8,
			expected: 		[]int{8, -1, -100, 3, 5, 10, 7},
		},
		{
			nums: 			[]int{-1, -100, 3, 5, 10},
			numsToRotate: 	8,
			expected: 		[]int{3, 5, 10, -1, -100},
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("Testing with expected=%d", tt.expected), func(t *testing.T) {
            // Call the function
            rotate(tt.nums, tt.numsToRotate)

            // Check if k is correct
            for i := 0; i < len(tt.nums); i++ {
                if tt.nums[i] != tt.expected[i] {
                    t.Errorf("Expected %d at index %d, got %d", tt.expected[i], i, tt.nums[i])
                }
            }
        })
	}
}
