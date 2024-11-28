package main

import (
	"fmt"
	"testing"
)

func Test_candy(t *testing.T) {
	tests := []struct {
		ratings  []int
		expected int
	}{
		{
			ratings:  []int{1, 0, 2},
			expected: 5,
		},
		{
			ratings:  []int{1, 2, 2},
			expected: 4,
		},
		{
			ratings:  []int{100, 80, 70, 60, 70, 80, 90, 100, 90, 80, 70, 60, 60},
			expected: 35,
		},
		{
			ratings:  []int{6, 7, 6, 5, 4, 3, 2, 1, 0, 0, 0, 1, 0},
			expected: 42,
		},
		{
			ratings:  []int{20000, 20000, 16001, 16001, 16002, 16002, 16003, 16003, 16004, 16004, 16005, 16005, 16006, 16006, 16007, 16007, 16008, 16008, 16009, 16009, 16010, 16010, 16011, 16011, 16012, 16012, 16013, 16013, 16014, 16014, 16015, 16015, 16016, 16016, 16017, 16017, 16018, 16018, 16019, 16019, 16020, 16020, 16021, 16021, 16022, 16022, 16023, 16023, 16024, 16024},
			expected: 74,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("Testing with expected=%d", tt.expected), func(t *testing.T) {
			// Call the function
			k := candy(tt.ratings)

			// Check if k is correct
			if k != tt.expected {
				t.Errorf("Expected %d, got %d", tt.expected, k)
			}
		})
	}
}
