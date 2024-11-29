package main

import (
	"fmt"
	"testing"
)

func Test_romanToInt(t *testing.T) {
	tests := []struct {
		s  			string
		expected 	int
	}{
		{
			s: 			"III",
			expected: 	3,
		},
		{
			s: 			"LVIII",
			expected: 	58,
		},
		{
			s: 			"MCMXCIV",
			expected: 	1994,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("Testing with expected=%d", tt.expected), func(t *testing.T) {
			// Call the function
			k := romanToInt(tt.s)

			// Check if k is correct
			if k != tt.expected {
				t.Errorf("Expected %d, got %d", tt.expected, k)
			}
		})
	}
}
