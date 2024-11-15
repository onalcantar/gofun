package main

import (
    "fmt"
    "sort"
    "testing"
)

func Test_removeDuplicates(t *testing.T) {
	tests := []struct {
		nums           []int
        val            int
        expectedNums   []int
	}{
		{
			nums: []int{1, 1, 1, 2, 2, 3},
			val: 5,
			expectedNums: []int{1, 1, 2, 2, 3},
		},
		{
			nums: []int{0, 0, 1, 1, 1, 1, 2, 3, 3},
			val: 7,
			expectedNums: []int{0, 0, 1, 1, 2, 3, 3},
		},
		{
			nums: []int{0},
			val: 1,
			expectedNums: []int{0},
		},
		{
			nums: []int{0, 0},
			val: 2,
			expectedNums: []int{0, 0},
		},
		{
			nums: []int{},
			val: 0,
			expectedNums: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("Testing with val=%d", tt.val), func(t *testing.T) {
            // Call the function
            k := removeDuplicates(tt.nums)

            // Sort both nums and expectedNums to compare in order
            sort.Ints(tt.nums[:k])
            sort.Ints(tt.expectedNums)

            // Check if k is correct
            if k != len(tt.expectedNums) {
                t.Errorf("Expected length %d, got %d", len(tt.expectedNums), k)
            }

            // Check if the result matches the expected values
            for i := 0; i < k; i++ {
                if tt.nums[i] != tt.expectedNums[i] {
                    t.Errorf("Expected %d at index %d, got %d", tt.expectedNums[i], i, tt.nums[i])
                }
            }
        })
	}
}
