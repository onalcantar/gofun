package main

import (
    "fmt"
    "sort"
    "testing"
)

func Test_removeElement(t *testing.T) {
    tests := []struct {
        nums           []int
        val            int
        expectedNums   []int
    }{
        {
            nums:         []int{3, 2, 2, 3},
            val:          3,
            expectedNums: []int{2, 2},
        },
        {
            nums:         []int{0, 1, 2, 2, 3, 0, 4, 2},
            val:          2,
            expectedNums: []int{0, 1, 3, 0, 4},
        },
        {
            nums:         []int{1, 2, 3, 4},
            val:          5,
            expectedNums: []int{1, 2, 3, 4}, // No removal
        },
        {
            nums:         []int{1, 2, 3, 4},
            val:          3,
            expectedNums: []int{1, 2, 4}, // 3 should be removed
        },
        {
            nums:         []int{2, 2, 2, 2},
            val:          2,
            expectedNums: []int{}, // All elements are removed
        },
        {
            nums:         []int{},
            val:          2,
            expectedNums: []int{}, // Empty array
        },
    }

    for _, tt := range tests {
        t.Run(fmt.Sprintf("Testing with val=%d", tt.val), func(t *testing.T) {
            // Call the function
            k := removeElement(tt.nums, tt.val)

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
