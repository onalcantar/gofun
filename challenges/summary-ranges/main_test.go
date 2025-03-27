package main

import (
	"fmt"
	"testing"
)

func Test_summaryRanges(t *testing.T) {
	tests := []struct {
		nums 	 	[]int
        expected   	[]string
	}{
		{
			nums: 		[]int{0, 1, 2, 4, 5, 7},
			expected: 	[]string{"0->2", "4->5", "7"},
		},
		{
			nums: 		[]int{0, 2, 3, 4, 6, 8, 9},
			expected: 	[]string{"0", "2->4", "6", "8->9"},
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("Testing with expected=%v", tt.expected), func(t *testing.T) {
            // Call the function
            k := summaryRanges(tt.nums)

            // Check if k is correct
			for i := range k {
                if k[i] != tt.expected[i] {
                    t.Errorf("Expected %s at index %d, got %s", tt.expected[i], i, k[i])
                }
            }
        })
	}
}
