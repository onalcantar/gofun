package main

import (
	"fmt"
	"testing"
)

func Test_canCompleteCircuit(t *testing.T) {
	tests := []struct {
		gas           []int
		cost          []int
        expected      int
	}{
		{
			gas: 		[]int{1, 2, 3, 4, 5},
			cost:		[]int{3, 4, 5, 1, 2},
			expected: 	3,
		},
		{
			gas: 		[]int{2, 3, 4},
			cost:		[]int{3, 4, 3},
			expected: 	-1,
		},
		{
			gas: 		[]int{4, 3, 4},
			cost:		[]int{3, 4, 3},
			expected: 	0,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("Testing with expected=%d", tt.expected), func(t *testing.T) {
            // Call the function
            k := canCompleteCircuit(tt.gas, tt.cost)

            // Check if k is correct
			if k != tt.expected {
				t.Errorf("Expected %d, got %d", k, tt.expected)
			}
        })
	}
}
